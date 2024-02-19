// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjsadow/msgraph.go/jsonx"
)

// CloudAppSecurityProfileRequestBuilder is request builder for CloudAppSecurityProfile
type CloudAppSecurityProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudAppSecurityProfileRequest
func (b *CloudAppSecurityProfileRequestBuilder) Request() *CloudAppSecurityProfileRequest {
	return &CloudAppSecurityProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudAppSecurityProfileRequest is request for CloudAppSecurityProfile
type CloudAppSecurityProfileRequest struct{ BaseRequest }

// Get performs GET request for CloudAppSecurityProfile
func (r *CloudAppSecurityProfileRequest) Get(ctx context.Context) (resObj *CloudAppSecurityProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudAppSecurityProfile
func (r *CloudAppSecurityProfileRequest) Update(ctx context.Context, reqObj *CloudAppSecurityProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudAppSecurityProfile
func (r *CloudAppSecurityProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CloudCommunicationsRequestBuilder is request builder for CloudCommunications
type CloudCommunicationsRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudCommunicationsRequest
func (b *CloudCommunicationsRequestBuilder) Request() *CloudCommunicationsRequest {
	return &CloudCommunicationsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudCommunicationsRequest is request for CloudCommunications
type CloudCommunicationsRequest struct{ BaseRequest }

// Get performs GET request for CloudCommunications
func (r *CloudCommunicationsRequest) Get(ctx context.Context) (resObj *CloudCommunications, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudCommunications
func (r *CloudCommunicationsRequest) Update(ctx context.Context, reqObj *CloudCommunications) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudCommunications
func (r *CloudCommunicationsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type CloudCommunicationsGetPresencesByUserIDRequestBuilder struct{ BaseRequestBuilder }

// GetPresencesByUserID action undocumented
func (b *CloudCommunicationsRequestBuilder) GetPresencesByUserID(reqObj *CloudCommunicationsGetPresencesByUserIDRequestParameter) *CloudCommunicationsGetPresencesByUserIDRequestBuilder {
	bb := &CloudCommunicationsGetPresencesByUserIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getPresencesByUserId"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CloudCommunicationsGetPresencesByUserIDRequest struct{ BaseRequest }

//
func (b *CloudCommunicationsGetPresencesByUserIDRequestBuilder) Request() *CloudCommunicationsGetPresencesByUserIDRequest {
	return &CloudCommunicationsGetPresencesByUserIDRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CloudCommunicationsGetPresencesByUserIDRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Presence, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Presence
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Presence
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *CloudCommunicationsGetPresencesByUserIDRequest) PostN(ctx context.Context, n int) ([]Presence, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *CloudCommunicationsGetPresencesByUserIDRequest) Post(ctx context.Context) ([]Presence, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}
