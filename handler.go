var Handlers = map[string]func([]Value) Value

//Create the command
func ping(args []Value) Value{
	return Value{typ: "string", str:"PONG"}
}
 //Add the command to the handler
var Handlers = map[string]func([]Value) Value{
	"PING": ping
}