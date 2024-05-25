// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Role - The role of the messages author, in this case `system`.
type Role string

const (
	RoleSystem Role = "system"
)

func (e Role) ToPointer() *Role {
	return &e
}
func (e *Role) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system":
		*e = Role(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Role: %v", v)
	}
}

type ChatCompletionRequestSystemMessage struct {
	// The contents of the system message.
	Content string `json:"content"`
	// The role of the messages author, in this case `system`.
	Role Role `json:"role"`
	// An optional name for the participant. Provides the model information to differentiate between participants of the same role.
	Name *string `json:"name,omitempty"`
}

func (o *ChatCompletionRequestSystemMessage) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *ChatCompletionRequestSystemMessage) GetRole() Role {
	if o == nil {
		return Role("")
	}
	return o.Role
}

func (o *ChatCompletionRequestSystemMessage) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
