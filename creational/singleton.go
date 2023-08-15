package creational

import "sync"

type Connection struct {
	Url  string
	Port string
}

var connection *Connection
var syncConnection *Connection

// With mutex
var lock = &sync.Mutex{}

func getInstance() *Connection {
	if connection != nil {
		println("Instance already there!")
	} else {
		println("Creating instance!")
		lock.Lock()
		defer lock.Unlock()
		connection = &Connection{}
	}
	return connection
}

// With sync.Once
var once sync.Once

func getInstanceOnce() *Connection {
	if connection == nil {
		println("[Sync Once] Creating instance!")
		once.Do(
			func() {
				syncConnection = &Connection{}
			},
		)
	} else {
		println("[Sync Once] Instance already there!")
	}
	return syncConnection
}
