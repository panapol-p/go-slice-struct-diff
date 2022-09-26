package listener

import (
	"encoding/json"
	"reflect"
	"sync"
)

type Listener[T any] struct {
	CurrentValue  map[string]Value[T]
	NewValue      map[string]Value[T]
	EventCallback func([]Events[T])
	CurrentEvent  []Events[T]
	mu            sync.Mutex
}

type EventState string

type DataStruct interface{}
type Value[T any] struct {
	Hash string
	Data T
}

var EventStateAdded EventState = "added"
var EventStateUpdated EventState = "updated"
var EventStateDeleted EventState = "deleted"

type Events[T any] struct {
	ID    string
	State EventState
	Data  T
}

func NewListener[T any]() *Listener[T] {
	return &Listener[T]{}
}

func (l *Listener[T]) SetCallback(f func([]Events[T])) {
	l.EventCallback = f
}

func (l *Listener[T]) AddNewValue(a []T) []Events[T] {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.NewValue = l.convertToMap(a)
	return l.compareMap()
}

func (l *Listener[T]) convertToMap(s []T) map[string]Value[T] {
	m := make(map[string]Value[T])

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
		m[uniqueID] = Value[T]{
			Hash: string(b),
			Data: s[ii],
		}
	}
	return m
}

func (l *Listener[T]) compareMap() []Events[T] {
	var es []Events[T]

	for key := range l.NewValue {
		var e Events[T]
		if value, ok := l.CurrentValue[key]; ok {
			if value.Hash != l.NewValue[key].Hash {
				e.ID = key
				e.State = EventStateUpdated
				e.Data = l.NewValue[key].Data
				es = append(es, e)
			}
		} else {
			e.ID = key
			e.State = EventStateAdded
			e.Data = l.NewValue[key].Data
			es = append(es, e)
		}
	}

	for key := range l.CurrentValue {
		var e Events[T]
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
