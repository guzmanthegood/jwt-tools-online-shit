package apix

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

// NewCachedClient returns new TravelgateX GateWay API cached client
func NewCachedClient(client Client) Client {
	c := cache.New(5*time.Minute, 10*time.Minute)

	return cachedClient{
		client: client,
		cache:  c,
	}
}

type cachedClient struct {
	client Client
	cache  *cache.Cache
}

func (c cachedClient) GetGroup(code string) (*Group, error) {
	cached, found := c.cache.Get(code)
	if found {
		return cached.(*Group), nil
	}

	gr, err := c.client.GetGroup(code)
	c.cache.Set(code, gr, cache.DefaultExpiration)
	return gr, err
}
