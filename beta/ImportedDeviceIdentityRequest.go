// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ImportedDeviceIdentityRequestBuilder is request builder for ImportedDeviceIdentity
type ImportedDeviceIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ImportedDeviceIdentityRequest
func (b *ImportedDeviceIdentityRequestBuilder) Request() *ImportedDeviceIdentityRequest {
	return &ImportedDeviceIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ImportedDeviceIdentityRequest is request for ImportedDeviceIdentity
type ImportedDeviceIdentityRequest struct{ BaseRequest }

// Do performs HTTP request for ImportedDeviceIdentity
func (r *ImportedDeviceIdentityRequest) Do(method, path string, reqObj interface{}) (resObj *ImportedDeviceIdentity, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ImportedDeviceIdentity
func (r *ImportedDeviceIdentityRequest) Get() (*ImportedDeviceIdentity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ImportedDeviceIdentity
func (r *ImportedDeviceIdentityRequest) Update(reqObj *ImportedDeviceIdentity) (*ImportedDeviceIdentity, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ImportedDeviceIdentity
func (r *ImportedDeviceIdentityRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}