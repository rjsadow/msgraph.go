// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjsadow/msgraph.go/jsonx"
)

// UserFlows returns request builder for IdentityUserFlow collection
func (b *IdentityContainerRequestBuilder) UserFlows() *IdentityContainerUserFlowsCollectionRequestBuilder {
	bb := &IdentityContainerUserFlowsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userFlows"
	return bb
}

// IdentityContainerUserFlowsCollectionRequestBuilder is request builder for IdentityUserFlow collection
type IdentityContainerUserFlowsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IdentityUserFlow collection
func (b *IdentityContainerUserFlowsCollectionRequestBuilder) Request() *IdentityContainerUserFlowsCollectionRequest {
	return &IdentityContainerUserFlowsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IdentityUserFlow item
func (b *IdentityContainerUserFlowsCollectionRequestBuilder) ID(id string) *IdentityUserFlowRequestBuilder {
	bb := &IdentityUserFlowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IdentityContainerUserFlowsCollectionRequest is request for IdentityUserFlow collection
type IdentityContainerUserFlowsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IdentityUserFlow collection
func (r *IdentityContainerUserFlowsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IdentityUserFlow, error) {
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
	var values []IdentityUserFlow
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
			value  []IdentityUserFlow
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

// GetN performs GET request for IdentityUserFlow collection, max N pages
func (r *IdentityContainerUserFlowsCollectionRequest) GetN(ctx context.Context, n int) ([]IdentityUserFlow, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IdentityUserFlow collection
func (r *IdentityContainerUserFlowsCollectionRequest) Get(ctx context.Context) ([]IdentityUserFlow, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IdentityUserFlow collection
func (r *IdentityContainerUserFlowsCollectionRequest) Add(ctx context.Context, reqObj *IdentityUserFlow) (resObj *IdentityUserFlow, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// EntitlementManagement is navigation property
func (b *IdentityGovernanceRequestBuilder) EntitlementManagement() *EntitlementManagementRequestBuilder {
	bb := &EntitlementManagementRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/entitlementManagement"
	return bb
}

// ImpactedUser is navigation property
func (b *IdentityRiskEventRequestBuilder) ImpactedUser() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/impactedUser"
	return bb
}
