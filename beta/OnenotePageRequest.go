// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OnenotePageRequestBuilder is request builder for OnenotePage
type OnenotePageRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenotePageRequest
func (b *OnenotePageRequestBuilder) Request() *OnenotePageRequest {
	return &OnenotePageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenotePageRequest is request for OnenotePage
type OnenotePageRequest struct{ BaseRequest }

// Do performs HTTP request for OnenotePage
func (r *OnenotePageRequest) Do(method, path string, reqObj interface{}) (resObj *OnenotePage, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OnenotePage
func (r *OnenotePageRequest) Get() (*OnenotePage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OnenotePage
func (r *OnenotePageRequest) Update(reqObj *OnenotePage) (*OnenotePage, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OnenotePage
func (r *OnenotePageRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ParentNotebook is navigation property
func (b *OnenotePageRequestBuilder) ParentNotebook() *NotebookRequestBuilder {
	bb := &NotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentNotebook"
	return bb
}

// ParentSection is navigation property
func (b *OnenotePageRequestBuilder) ParentSection() *OnenoteSectionRequestBuilder {
	bb := &OnenoteSectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentSection"
	return bb
}