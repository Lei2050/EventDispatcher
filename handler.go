package EventDispatcher

type EventArg map[string]any

func NewEventArg() EventArg {
	return make(map[string]any)
}

func (e EventArg) Add(key string, a any) EventArg {
	e[key] = a
	return e
}

func (e EventArg) GetValue(key string) any {
	return e[key]
}

func (e EventArg) GetInt(key string) int {
	p := e[key]
	if p == nil {
		return 0
	}
	switch v := p.(type) {
	case int:
		return v
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	}
	return 0
}

func (e EventArg) GetUint32(key string) uint32 {
	p := e[key]
	if p == nil {
		return 0
	}
	switch v := p.(type) {
	case int:
		return uint32(v)
	case uint32:
		return v
	case uint64:
		return uint32(v)
	}
	return 0
}

func (e EventArg) GetUint64(key string) uint64 {
	p := e[key]
	if p == nil {
		return 0
	}
	switch v := p.(type) {
	case int:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	}
	return 0
}

func (e EventArg) GetStr(key string) string {
	p := e[key]
	if p == nil {
		return ""
	}
	v, ok := p.(string)
	if !ok {
		return ""
	}
	return v
}

func (e EventArg) Get(key string) any {
	if v, exist := e[key]; exist {
		return v
	}
	return nil
}

type EventFunc func(a EventArg)
