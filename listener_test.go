package listener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type S struct {
	F0 string `listener:"id"`
	F1 string
	F2 string
}

func TestNewListener(t *testing.T) {
	l := NewListener[S]()
	assert.NotNil(t, l)
}

func TestListener_SetCallback(t *testing.T) {
	l := NewListener[S]()
	f := func(e []Events) {}
	l.SetCallback(f)
	assert.NotNil(t, l.EventCallback)
}

func TestListener_AddNewValue(t *testing.T) {
	l := new(Listener[S])

	s := []S{
		{F0: "1", F1: "test2", F2: ""},
		{F0: "2", F1: "test2", F2: ""},
	}
	l.AddNewValue(s)
	expected := map[string]string{
		"1": "{\"F0\":\"1\",\"F1\":\"test2\",\"F2\":\"\"}",
		"2": "{\"F0\":\"2\",\"F1\":\"test2\",\"F2\":\"\"}",
	}
	assert.Equal(t, expected, l.NewValue)
}

func TestListener_convertToMap(t *testing.T) {
	l := new(Listener[S])

	s := []S{
		{F0: "1", F1: "test2", F2: ""},
		{F0: "2", F1: "test2", F2: ""},
	}
	actual := l.convertToMap(s)
	expected := map[string]string{
		"1": "{\"F0\":\"1\",\"F1\":\"test2\",\"F2\":\"\"}",
		"2": "{\"F0\":\"2\",\"F1\":\"test2\",\"F2\":\"\"}",
	}
	assert.Equal(t, expected, actual)
}

func TestListener_compareMap(t *testing.T) {
	l := new(Listener[S])
	l.CurrentValue = map[string]string{
		"2": "2",
		"3": "6",
		"4": "4",
		"5": "5",
	}
	l.NewValue = map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
	}

	e := l.compareMap()
	expected := []Events{
		{State: EventStateAdded, ID: "1"},
		{State: EventStateUpdated, ID: "3"},
		{State: EventStateDeleted, ID: "5"},
	}
	assert.Equal(t, expected, e)

	// test compare with callback
	f := func(e []Events) {}
	l.SetCallback(f)
	e = l.compareMap()
	//no change event
	assert.Nil(t, e)
}
