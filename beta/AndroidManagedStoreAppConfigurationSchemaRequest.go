// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidManagedStoreAppConfigurationSchemaRequestBuilder is request builder for AndroidManagedStoreAppConfigurationSchema
type AndroidManagedStoreAppConfigurationSchemaRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidManagedStoreAppConfigurationSchemaRequest
func (b *AndroidManagedStoreAppConfigurationSchemaRequestBuilder) Request() *AndroidManagedStoreAppConfigurationSchemaRequest {
	return &AndroidManagedStoreAppConfigurationSchemaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidManagedStoreAppConfigurationSchemaRequest is request for AndroidManagedStoreAppConfigurationSchema
type AndroidManagedStoreAppConfigurationSchemaRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidManagedStoreAppConfigurationSchema
func (r *AndroidManagedStoreAppConfigurationSchemaRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidManagedStoreAppConfigurationSchema, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidManagedStoreAppConfigurationSchema
func (r *AndroidManagedStoreAppConfigurationSchemaRequest) Get() (*AndroidManagedStoreAppConfigurationSchema, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidManagedStoreAppConfigurationSchema
func (r *AndroidManagedStoreAppConfigurationSchemaRequest) Update(reqObj *AndroidManagedStoreAppConfigurationSchema) (*AndroidManagedStoreAppConfigurationSchema, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidManagedStoreAppConfigurationSchema
func (r *AndroidManagedStoreAppConfigurationSchemaRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}