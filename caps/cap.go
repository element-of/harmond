package caps

import "sync"

// Cap represents an individual IRC server capability
type Cap struct {
	Name  string
	Value string
}

// Clone creates a copy of this Cap
func (c Cap) Clone() Cap {
	return Cap{
		Name:  c.Name,
		Value: c.Value,
	}
}

//
func (c Cap) CloneWithValue(val string) Cap {
	return Cap{
		Name:  c.Name,
		Value: val,
	}
}

func (c Cap) IRCv31Show() string {
	return c.Name
}

func (c Cap) IRCv32Show() string {
	return c.Name + "=" + c.Value
}

// CapSet represents a set of capabilities that a Session enabled.
type CapSet struct {
	sync.Mutex
	caps []Cap
}

func (cs *CapSet) HasCap(name string) (bool, string) {
	cs.Lock()
	defer cs.Unlock()

	for _, c := range cs.caps {
		if c.Name == name {
			return true, c.Value
		}
	}

	return false, ""
}

func (cs *CapSet) AddCap(c Cap) {
	cs.Lock()
	defer cs.Unlock()

	has, _ := cs.HasCap(c.Name)
	if !has {
		cs.caps = append(cs.caps, c)
	}
}

type CapRegistry struct {
	sync.Mutex
	caps map[string]Cap
}

func (cr *CapRegistry) HasCap(name string) (Cap, bool) {
	cr.Lock()
	defer cr.Unlock()

	cap, ok := cr.caps[name]
	return cap.Clone(), ok
}
