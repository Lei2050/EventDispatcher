package event

import (
	"testing"
)

func TestEventMgr(t *testing.T) {
	const (
		EVENT_ID_LEVEL_UP int = iota
		EVENT_ID_CONSUME_MONEY
		EVENT_ID_CONSUME_ITEM
	)

	em := NewEventMgr()
	em.RegisterEvent(EVENT_ID_LEVEL_UP, func(a EventArg) {
		var (
			uid uint64 = a.GetUint64("uid")
			lv  uint32 = a.GetUint32("lv")
		)
		t.Logf("levelup uid:%d, lv:%d", uid, lv)
	})
	em.RegisterEvent(EVENT_ID_CONSUME_MONEY, func(a EventArg) {
		var (
			uid  uint64 = a.GetUint64("uid")
			cost uint32 = a.GetUint32("cost")
		)
		t.Logf("consume money uid:%d, lv:%d", uid, cost)
	})
	em.RegisterEvent(EVENT_ID_CONSUME_ITEM, func(a EventArg) {
		var (
			uid    uint64 = a.GetUint64("uid")
			itemId uint32 = a.GetUint32("item_id")
			num    uint32 = a.GetUint32("num")
		)
		t.Logf("consume item uid:%d, item_id:%d, lv:%d", uid, itemId, num)
	})

	em.SendEvent(EVENT_ID_LEVEL_UP, map[string]interface{}{"uid": 12138, "lv": 99})
	em.SendEvent(EVENT_ID_CONSUME_MONEY, map[string]interface{}{"uid": uint64(12138), "cost": uint32(1000000)})
	//em.SendEvent2(EVENT_ID_CONSUME_ITEM, &ConsumeItemParam{Uid: 12138, ItemId: 110, Num: 5})
	em.SendEvent2(EVENT_ID_CONSUME_ITEM, NewEventArg().Add("uid", 12138).Add("item_id", 110).Add("num", 5))

	RegisterEventS(EVENT_ID_LEVEL_UP, func(a EventArg) {
		var (
			uid uint64 = a.GetUint64("uid")
			lv  uint32 = a.GetUint32("lv")
		)
		t.Logf("levelup uid:%d, lv:%d", uid, lv)
	})
	RegisterEventS(EVENT_ID_CONSUME_MONEY, func(a EventArg) {
		var (
			uid  uint64 = a.GetUint64("uid")
			cost uint32 = a.GetUint32("cost")
		)
		t.Logf("consume money uid:%d, lv:%d", uid, cost)
	})
	RegisterEventS(EVENT_ID_CONSUME_ITEM, func(a EventArg) {
		var (
			uid    uint64 = a.GetUint64("uid")
			itemId uint32 = a.GetUint32("item_id")
			num    uint32 = a.GetUint32("num")
		)
		t.Logf("consume item uid:%d, item_id:%d, lv:%d", uid, itemId, num)
	})

	SendEventS(EVENT_ID_LEVEL_UP, map[string]interface{}{"uid": 12138, "lv": 99})
	SendEventS(EVENT_ID_CONSUME_MONEY, map[string]interface{}{"uid": uint64(12138), "cost": uint32(1000000)})
	SendEventS(EVENT_ID_CONSUME_ITEM, NewEventArg().Add("uid", 12138).Add("item_id", 110).Add("num", 5))
}
