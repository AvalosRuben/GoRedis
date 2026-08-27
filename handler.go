package main

import "sync"

//Create the command
func ping(args []Value) Value{
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}
	return Value{typ: "string", str: args[0].bulk}
}
 //Add the command to the handler
var Handlers = map[string]func([]Value) Value{
	"PING": ping,
}

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

func set(args []Value) Value{
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].bulk
	value := args[1].bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return Value{typ:"string", str:"Ok"}
}