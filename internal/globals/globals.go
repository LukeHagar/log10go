// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package globals

type Globals struct {
	XLog10Organization *string `header:"style=simple,explode=false,name=X-Log10-Organization"`
}

func (o *Globals) GetXLog10Organization() *string {
	if o == nil {
		return nil
	}
	return o.XLog10Organization
}
