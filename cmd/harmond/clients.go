package main

import cmap "github.com/orcaman/concurrent-map"

// ClientSet is the set of clients known to this IRC server
type ClientSet struct {
	byNick cmap.ConcurrentMap
	byID   cmap.ConcurrentMap
}

func newClientSet() *ClientSet {
	return &ClientSet{
		byNick: cmap.New(),
		byID:   cmap.New(),
	}
}

// Add adds a new Client to the pool
func (c *ClientSet) Add(cl *Client) bool {
	sn := c.byNick.SetIfAbsent(NickSort(cl.Nick), cl)
	su := c.byID.SetIfAbsent(cl.ID, cl)
	return sn == su
}

// GetNick returns a client by nickname.
func (c *ClientSet) GetNick(nick string) (*Client, bool) {
	data, ok := c.byNick.Get(NickSort(nick))
	if !ok {
		return nil, false
	}

	cl, ok := data.(*Client)
	return cl, ok
}

// GetID returns a client by their unique ID.
func (c *ClientSet) GetID(id string) (*Client, bool) {
	data, ok := c.byID.Get(id)
	if !ok {
		return nil, false
	}

	cl, ok := data.(*Client)
	return cl, ok
}

// Del removes a client from the client set.
func (c *ClientSet) Del(cl *Client) {
	c.byID.Remove(cl.ID)
	c.byNick.Remove(NickSort(cl.Nick))
}

// ChangeNick updates a client's nickname in the client set.
func (c *ClientSet) ChangeNick(cl *Client, oldnick string) {
	c.byNick.Remove(NickSort(oldnick))
	c.byNick.Set(NickSort(cl.Nick), cl)
}
