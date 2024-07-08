package EventDispatcher

import "reflect"

type EventMgr struct {
	efs map[int][]EventFunc
}

func NewEventMgr() *EventMgr {
	return &EventMgr{
		efs: make(map[int][]EventFunc),
	}
}

func (em *EventMgr) RegisterEvent(id int, f EventFunc) *EventMgr {
	em.efs[id] = append(em.efs[id], f)
	return em
}

func (em *EventMgr) UnregisterEvent(id int, f EventFunc) *EventMgr {
	handlers := em.efs[id]
	var handlerSize = len(handlers)
	if handlerSize <= 0 {
		return em
	}

	pf := reflect.ValueOf(f).Pointer()
	idx := -1
	for k, v := range handlers {
		if reflect.ValueOf(v).Pointer() == pf {
			idx = k
			break
		}
	}
	if idx != -1 {
		if idx != handlerSize-1 {
			handlers[idx], handlers[handlerSize-1] = handlers[handlerSize-1], nil
		}
		handlers = handlers[:handlerSize-1]
		em.efs[id] = handlers
	}
	return em
}

func (em *EventMgr) SendEvent(id int, args map[string]interface{}) {
	em.SendEvent2(id, EventArg(args))
}

// 与SendEvent相比，就是传参数方式不一样
func (em *EventMgr) SendEvent2(id int, a EventArg) {
	for _, f := range em.efs[id] {
		f(a)
	}
}

var g_EventMgr *EventMgr = NewEventMgr()

func RegisterEvent(id int, f EventFunc) *EventMgr {
	return g_EventMgr.RegisterEvent(id, f)
}

func UnregisterEvent(id int, f EventFunc) *EventMgr {
	return g_EventMgr.UnregisterEvent(id, f)
}

func SendEvent(id int, args map[string]interface{}) {
	g_EventMgr.SendEvent(id, args)
}

func SendEvent2(id int, a EventArg) {
	g_EventMgr.SendEvent2(id, a)
}
