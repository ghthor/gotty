package server

import (
	"github.com/ghthor/gotty/v2/webtty"

	"github.com/gorilla/websocket"
)

// Slave is webtty.Slave with some additional methods.
type Slave interface {
	webtty.Slave

	Close() error
}

type Factory interface {
	Name() string
	New(params map[string][]string, conn *websocket.Conn) (Slave, error)
}
