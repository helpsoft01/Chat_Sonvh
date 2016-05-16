package model

import (
	"goredis-master"
	"time"
	"fmt"
)

type Client struct {
	r *goredis.Redis
}

var dial = goredis.DialConfig{
	Network: goredis.DefaultNetwork,
	Address:"127.0.0.1" + goredis.DefaultAddress,
	Database:1,
	MaxIdle:1,
	Password:"",
	Timeout:5 * time.Second }
var ConfigDBURLConnect = "tcp://auth:%s@%s/%d?timeout=%s&maxidle=%d"

var client = new(Client)

func (c *Client)Init() {
	//tcp://auth:password@127.0.0.1:6379/0?timeout=10s&maxidle=1

	if client == nil || client.r == nil {
		connectString := fmt.Sprintf(ConfigDBURLConnect, dial.Password, dial.Address, dial.Database, dial.Timeout.String(), dial.MaxIdle)
		_client, _ := goredis.DialURL(connectString)
		client.r = _client
		//defer client.r.ClosePool()
	}
}
func (c *Client)GetClient() *goredis.Redis {
	return client.r
}
