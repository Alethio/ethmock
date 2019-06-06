package server

var methodNotFoundResponse = `
{  
	"jsonrpc":"2.0",
	"id":"1",
	"error":{  
	   "code":-32601,
	   "message":"The method %s does not exist/is not available"
	}
 }`
