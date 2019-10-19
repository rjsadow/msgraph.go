// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PaymentTermRequestBuilder is request builder for PaymentTerm
type PaymentTermRequestBuilder struct{ BaseRequestBuilder }

// Request returns PaymentTermRequest
func (b *PaymentTermRequestBuilder) Request() *PaymentTermRequest {
	return &PaymentTermRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PaymentTermRequest is request for PaymentTerm
type PaymentTermRequest struct{ BaseRequest }

// Do performs HTTP request for PaymentTerm
func (r *PaymentTermRequest) Do(method, path string, reqObj interface{}) (resObj *PaymentTerm, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for PaymentTerm
func (r *PaymentTermRequest) Get() (*PaymentTerm, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for PaymentTerm
func (r *PaymentTermRequest) Update(reqObj *PaymentTerm) (*PaymentTerm, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for PaymentTerm
func (r *PaymentTermRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}