// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/log10-io/log10go/models/components"
)

type ListFeedbackTasksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Tasks []components.Task
}

func (o *ListFeedbackTasksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListFeedbackTasksResponse) GetTasks() []components.Task {
	if o == nil {
		return nil
	}
	return o.Tasks
}
