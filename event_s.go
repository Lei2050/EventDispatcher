package EventDispatcher

import "sync"

//没什么区别，就是线程安全版本
type EventMgrS struct {
	efs map[int][]EventFunc
	sync.RWMutex
}

func NewEventMgrS() *EventMgrS {
	return &EventMgrS{
		efs: make(map[int][]EventFunc),
	}
}

func (em *EventMgrS) RegisterEvent(id int, f EventFunc) *EventMgrS {
	em.Lock()
	em.efs[id] = append(em.efs[id], f)
	em.Unlock()
	return em
}

func (em *EventMgrS) SendEvent(id int, args map[string]interface{}) {
	em.SendEvent2(id, EventArg(args))
}

//与SendEvent相比，就是传参数方式不一样
func (em *EventMgrS) SendEvent2(id int, a EventArg) {
	em.RLock()
	efs := em.efs[id]
	em.RUnlock()

	for _, f := range efs {
		f(a)
	}
}

var g_EventMgrS *EventMgrS = NewEventMgrS()

func RegisterEventS(id int, f EventFunc) *EventMgrS {
	return g_EventMgrS.RegisterEvent(id, f)
}

func SendEventS(id int, args map[string]interface{}) {
	g_EventMgrS.SendEvent(id, args)
}

func SendEvent2S(id int, a EventArg) {
	g_EventMgrS.SendEvent2(id, a)
}
