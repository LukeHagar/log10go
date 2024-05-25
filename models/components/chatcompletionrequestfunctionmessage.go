// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ChatCompletionRequestFunctionMessageRole - The role of the messages author, in this case `function`.
type ChatCompletionRequestFunctionMessageRole string

const (
	ChatCompletionRequestFunctionMessageRoleFunction ChatCompletionRequestFunctionMessageRole = "function"
)

func (e ChatCompletionRequestFunctionMessageRole) ToPointer() *ChatCompletionRequestFunctionMessageRole {
	return &e
}
func (e *ChatCompletionRequestFunctionMessageRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "function":
		*e = ChatCompletionRequestFunctionMessageRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionRequestFunctionMessageRole: %v", v)
	}
}

// ChatCompletionRequestFunctionMessage
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type ChatCompletionRequestFunctionMessage struct {
	// The role of the messages author, in this case `function`.
	Role ChatCompletionRequestFunctionMessageRole `json:"role"`
	// The contents of the function message.
	Content *string `json:"content"`
	// The name of the function to call.
	Name string `json:"name"`
}

func (o *ChatCompletionRequestFunctionMessage) GetRole() ChatCompletionRequestFunctionMessageRole {
	if o == nil {
		return ChatCompletionRequestFunctionMessageRole("")
	}
	return o.Role
}

func (o *ChatCompletionRequestFunctionMessage) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *ChatCompletionRequestFunctionMessage) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
