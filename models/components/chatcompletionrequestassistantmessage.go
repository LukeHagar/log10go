// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ChatCompletionRequestAssistantMessageRole - The role of the messages author, in this case `assistant`.
type ChatCompletionRequestAssistantMessageRole string

const (
	ChatCompletionRequestAssistantMessageRoleAssistant ChatCompletionRequestAssistantMessageRole = "assistant"
)

func (e ChatCompletionRequestAssistantMessageRole) ToPointer() *ChatCompletionRequestAssistantMessageRole {
	return &e
}
func (e *ChatCompletionRequestAssistantMessageRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "assistant":
		*e = ChatCompletionRequestAssistantMessageRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionRequestAssistantMessageRole: %v", v)
	}
}

// FunctionCall - Deprecated and replaced by `tool_calls`. The name and arguments of a function that should be called, as generated by the model.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type FunctionCall struct {
	// The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function.
	Arguments string `json:"arguments"`
	// The name of the function to call.
	Name string `json:"name"`
}

func (o *FunctionCall) GetArguments() string {
	if o == nil {
		return ""
	}
	return o.Arguments
}

func (o *FunctionCall) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type ChatCompletionRequestAssistantMessage struct {
	// The contents of the assistant message. Required unless `tool_calls` or `function_call` is specified.
	//
	Content *string `json:"content,omitempty"`
	// The role of the messages author, in this case `assistant`.
	Role ChatCompletionRequestAssistantMessageRole `json:"role"`
	// An optional name for the participant. Provides the model information to differentiate between participants of the same role.
	Name *string `json:"name,omitempty"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls []ChatCompletionMessageToolCall `json:"tool_calls,omitempty"`
	// Deprecated and replaced by `tool_calls`. The name and arguments of a function that should be called, as generated by the model.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
}

func (o *ChatCompletionRequestAssistantMessage) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *ChatCompletionRequestAssistantMessage) GetRole() ChatCompletionRequestAssistantMessageRole {
	if o == nil {
		return ChatCompletionRequestAssistantMessageRole("")
	}
	return o.Role
}

func (o *ChatCompletionRequestAssistantMessage) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ChatCompletionRequestAssistantMessage) GetToolCalls() []ChatCompletionMessageToolCall {
	if o == nil {
		return nil
	}
	return o.ToolCalls
}

func (o *ChatCompletionRequestAssistantMessage) GetFunctionCall() *FunctionCall {
	if o == nil {
		return nil
	}
	return o.FunctionCall
}
