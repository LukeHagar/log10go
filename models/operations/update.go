// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/log10-io/log10go/models/components"
)

type UpdateGlobals struct {
	XLog10Organization string `header:"style=simple,explode=false,name=X-Log10-Organization"`
}

func (o *UpdateGlobals) GetXLog10Organization() string {
	if o == nil {
		return ""
	}
	return o.XLog10Organization
}

type UpdateRequest struct {
	// The completion id to update.
	CompletionID       string                `pathParam:"style=simple,explode=false,name=completionId"`
	XLog10Organization *string               `header:"style=simple,explode=false,name=X-Log10-Organization"`
	Completion         components.Completion `request:"mediaType=application/json"`
}

func (o *UpdateRequest) GetCompletionID() string {
	if o == nil {
		return ""
	}
	return o.CompletionID
}

func (o *UpdateRequest) GetXLog10Organization() *string {
	if o == nil {
		return nil
	}
	return o.XLog10Organization
}

func (o *UpdateRequest) GetCompletion() components.Completion {
	if o == nil {
		return components.Completion{}
	}
	return o.Completion
}

type UpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Completion *components.Completion
}

func (o *UpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateResponse) GetCompletion() *components.Completion {
	if o == nil {
		return nil
	}
	return o.Completion
}
