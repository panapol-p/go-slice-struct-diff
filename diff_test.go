package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type S struct {
	F0 string `diff:"id"`
	F1 string
	F2 string
}

func TestNewDiff(t *testing.T) {
	l := NewDiff[S]()
	assert.NotNil(t, l)
}

func TestDiff_SetCallback(t *testing.T) {
	l := NewDiff[S]()
	f := func(e []Events[S]) {}
	l.SetCallback(f)
	assert.NotNil(t, l.EventCallback)
}

func TestDiff_AddNewValue(t *testing.T) {
	l := new(Diff[S])

	s := []S{
		{F0: "1", F1: "test2", F2: ""},
		{F0: "2", F1: "test2", F2: ""},
	}
	l.AddNewValue(s)
	expected := map[string]Value[S]{
		"1": {
			Hash: "{\"F0\":\"1\",\"F1\":\"test2\",\"F2\":\"\"}",
			Data: S{F0: "1", F1: "test2", F2: ""},
		},
		"2": {
			Hash: "{\"F0\":\"2\",\"F1\":\"test2\",\"F2\":\"\"}",
			Data: S{F0: "2", F1: "test2", F2: ""},
		}}
	assert.Equal(t, expected, l.NewValue)
}

func TestDiff_convertToMap(t *testing.T) {
	l := new(Diff[S])

	s := []S{
		{F0: "1", F1: "test2", F2: ""},
		{F0: "2", F1: "test2", F2: ""},
	}
	actual := l.convertToMap(s)
	expected := map[string]Value[S]{
		"1": {
			Hash: "{\"F0\":\"1\",\"F1\":\"test2\",\"F2\":\"\"}",
			Data: S{F0: "1", F1: "test2", F2: ""},
		},
		"2": {
			Hash: "{\"F0\":\"2\",\"F1\":\"test2\",\"F2\":\"\"}",
			Data: S{F0: "2", F1: "test2", F2: ""},
		}}
	assert.Equal(t, expected, actual)
}

func TestDiff_compareMap(t *testing.T) {
	l := new(Diff[S])
	l.CurrentValue = map[string]Value[S]{
		"2": {"{\"F0\":\"2\"}", S{F0: "2"}},
		"3": {"{\"F0\":\"3\"}", S{F0: "3"}},
		"4": {"{\"F0\":\"4\"}", S{F0: "4"}},
		"5": {"{\"F0\":\"5\"}", S{F0: "5"}},
	}
	l.NewValue = map[string]Value[S]{
		"1": {"{\"F0\":\"1\"}", S{F0: "1"}},
		"2": {"{\"F0\":\"2\"}", S{F0: "2"}},
		"3": {"{\"F0\":\"3\",\"F2\": \"4\"}}", S{F0: "3", F2: "4"}},
		"4": {"{\"F0\":\"4\"}", S{F0: "4"}},
	}

	e := l.compareMap()
	expected := []Events[S]{
		{State: EventStateAdded, ID: "1", Data: S{F0: "1"}},
		{State: EventStateUpdated, ID: "3", Data: S{F0: "3", F2: "4"}},
		{State: EventStateDeleted, ID: "5", Data: S{}},
	}
	assert.Equal(t, expected, e)

	// test compare with callback (no even data)
	f := func(e []Events[S]) {}
	l.SetCallback(f)
	e = l.compareMap()
	//no change event
	assert.Nil(t, e)

	// test compare with callback (have even data)
	s := []S{
		{F0: "1", F1: "test2", F2: ""},
		{F0: "2", F1: "test2", F2: ""},
	}
	l.AddNewValue(s)
	expected = []Events[S]{
		{State: EventStateUpdated, ID: "1", Data: S{F0: "1", F1: "test2", F2: ""}},
		{State: EventStateUpdated, ID: "2", Data: S{F0: "2", F1: "test2", F2: ""}},
		{State: EventStateDeleted, ID: "3", Data: S{}},
		{State: EventStateDeleted, ID: "4", Data: S{}},
	}

	assert.Equal(t, expected, l.CurrentEvent)
}
