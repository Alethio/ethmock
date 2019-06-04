package types

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
)

// JSONRPCRequest standard JSONRPC 2 Request
type JSONRPCRequest struct {
	Version string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

// Path returns a folder path generated from method
func (j JSONRPCRequest) Path() string {
	// create a standard folder structure
	segments := []string{j.Method}

	for _, s := range j.Params {
		str := fmt.Sprintf("%v", s)
		segments = append(segments, str)
	}
	folder := filepath.Join(segments...)
	return folder
}

// NewJSONRPCRequest creates a new json rpc request object
func NewJSONRPCRequest(method string, args []interface{}) *JSONRPCRequest {
	id := strconv.FormatInt(rand.Int63(), 16)

	return &JSONRPCRequest{
		Method: method,
		Params: args,
		ID:     id,
	}
}

// JSONRPCResponse standard JSONRPC 2 Response
type JSONRPCResponse struct {
	Version string          `json:"jsonrpc"`
	ID      string          `json:"id,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *JSONRPCError   `json:"error,omitempty"`
}

// JSONRPCError optional error
type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}
