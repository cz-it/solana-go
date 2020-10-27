package solana

import (
	"github.com/ethereum/go-ethereum/rpc"
)

// Connection is a connection to API node
type Connection struct {
	client *rpc.Client
}

// NewConnection create a connection to endpoint
func NewConnection(endpoint, commitment string) *Connection {
	c := &Connection{}
	c.client, _ = rpc.Dial(endpoint)
	return c
}
