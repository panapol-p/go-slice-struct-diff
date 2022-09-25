package listener

import (
	"encoding/json"
	"reflect"
	"sync"
)

type Listener[T any] struct {
	CurrentValue  map[string]string
	NewValue      map[string]string
	EventCallback func([]Events)
	CurrentEvent  []Events
	mu            sync.Mutex
}

type EventState string

type DataStruct interface{}

var EventStateAdded EventState = "added"
var EventStateUpdated EventState = "updated"
var EventStateDeleted EventState = "deleted"

type Events struct {
	State EventState
	ID    string
}

func NewListener[T any]() *Listener[T] {
	return &Listener[T]{}
}

func (l *Listener[T]) SetCallback(f func([]Events)) {
	l.EventCallback = f
}

func (l *Listener[T]) AddNewValue(a []T) []Events {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.NewValue = l.convertToMap(a)
	return l.compareMap()
}

func (l *Listener[T]) convertToMap(s []T) map[string]string {
	m := make(map[string]string)

	var uniqueField []string
	for ii := range s {
		var uniqueID string
		if ii == 0 {
			st := reflect.TypeOf(s[ii])
			for i := 0; i < st.NumField(); i++ {
				field := st.Field(i)
				if listener, ok := field.Tag.Lookup("listener"); ok {
					if listener == "id" {
						uniqueField = append(uniqueField, field.Name)
					}
				}
			}
		}

		r := reflect.ValueOf(s[ii])
		for j := range uniqueField {
			f := reflect.Indirect(r).FieldByName(uniqueField[j])
			uniqueID = uniqueID + f.String()
		}

		b, _ := json.Marshal(s[ii])
		m[uniqueID] = string(b)
	}
	return m
}

func (l *Listener[T]) compareMap() []Events {
	var es []Events

	for key := range l.NewValue {
		var e Events
		if value, ok := l.CurrentValue[key]; ok {
			if value != l.NewValue[key] {
				e.ID = key
				e.State = EventStateUpdated
				es = append(es, e)
			}
		} else {
			e.ID = key
			e.State = EventStateAdded
			es = append(es, e)
		}
	}

	for key := range l.CurrentValue {
		var e Events
		if _, ok := l.NewValue[key]; !ok {
			e.ID = key
			e.State = EventStateDeleted
			es = append(es, e)
		}
	}

	if l.EventCallback != nil {
		l.EventCallback(es)
	}

	l.CurrentValue = l.NewValue
	return es
}
