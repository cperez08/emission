package emission

import (
	"fmt"
)

const DEFAULT_MAX_LISTENERS = 10

type listener func(...interface{})

type Event struct {
	listeners []listener
}

type Emitter struct {
	events       map[string]*Event
	maxListeners int
}

func (e *Emitter) AddListener(s string, l func(...interface{})) {
	l = listener(l)
	_, ok := e.events[s]
	if !ok {
		e.events[s] = &Event{[]listener{}}
	}
	e.events[s].listeners = append(e.events[s].listeners, l)
}

func (e *Emitter) On(s string, l func(...interface{})) {
	e.AddListener(s, l)
}

func (e *Emitter) RemoveListener(l func(...interface{})) {
	for _, x := range e.events {
		for i, y := range x.listeners {
			if fmt.Sprintf("%v", y) == fmt.Sprintf("%v", l) {
				x.listeners = append(x.listeners[:i], x.listeners[i+1:]...)
			}
		}
	}
}

func (e *Emitter) Off(l func(...interface{})) {
	e.RemoveListener(l)
}

func (e *Emitter) Once(s string, l func(...interface{})) {
	e.AddListener(s, l)
	e.AddListener(s, func(i ...interface{}) {
		e.RemoveListener(l)
	})
}

func (e *Emitter) Emit(s string, i ...interface{}) {
	if _, ok := e.events[s]; !ok {
		return
	}
	for _, l := range e.events[s].listeners {
		l(i...)
	}
}

func NewEmitter() *Emitter {
	return &Emitter{
		make(map[string]*Event),
		DEFAULT_MAX_LISTENERS,
	}
}
