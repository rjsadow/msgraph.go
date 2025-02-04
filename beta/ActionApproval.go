// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjsadow/msgraph.go/jsonx"
)

// BusinessFlows returns request builder for BusinessFlow collection
func (b *ApprovalWorkflowProviderRequestBuilder) BusinessFlows() *ApprovalWorkflowProviderBusinessFlowsCollectionRequestBuilder {
	bb := &ApprovalWorkflowProviderBusinessFlowsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/businessFlows"
	return bb
}

// ApprovalWorkflowProviderBusinessFlowsCollectionRequestBuilder is request builder for BusinessFlow collection
type ApprovalWorkflowProviderBusinessFlowsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BusinessFlow collection
func (b *ApprovalWorkflowProviderBusinessFlowsCollectionRequestBuilder) Request() *ApprovalWorkflowProviderBusinessFlowsCollectionRequest {
	return &ApprovalWorkflowProviderBusinessFlowsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BusinessFlow item
func (b *ApprovalWorkflowProviderBusinessFlowsCollectionRequestBuilder) ID(id string) *BusinessFlowRequestBuilder {
	bb := &BusinessFlowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApprovalWorkflowProviderBusinessFlowsCollectionRequest is request for BusinessFlow collection
type ApprovalWorkflowProviderBusinessFlowsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BusinessFlow collection
func (r *ApprovalWorkflowProviderBusinessFlowsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BusinessFlow, error) {
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
	var values []BusinessFlow
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
			value  []BusinessFlow
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

// GetN performs GET request for BusinessFlow collection, max N pages
func (r *ApprovalWorkflowProviderBusinessFlowsCollectionRequest) GetN(ctx context.Context, n int) ([]BusinessFlow, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BusinessFlow collection
func (r *ApprovalWorkflowProviderBusinessFlowsCollectionRequest) Get(ctx context.Context) ([]BusinessFlow, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BusinessFlow collection
func (r *ApprovalWorkflowProviderBusinessFlowsCollectionRequest) Add(ctx context.Context, reqObj *BusinessFlow) (resObj *BusinessFlow, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BusinessFlowsWithRequestsAwaitingMyDecision returns request builder for BusinessFlow collection
func (b *ApprovalWorkflowProviderRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecision() *ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequestBuilder {
	bb := &ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/businessFlowsWithRequestsAwaitingMyDecision"
	return bb
}

// ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequestBuilder is request builder for BusinessFlow collection
type ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BusinessFlow collection
func (b *ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequestBuilder) Request() *ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequest {
	return &ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BusinessFlow item
func (b *ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequestBuilder) ID(id string) *BusinessFlowRequestBuilder {
	bb := &BusinessFlowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequest is request for BusinessFlow collection
type ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BusinessFlow collection
func (r *ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BusinessFlow, error) {
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
	var values []BusinessFlow
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
			value  []BusinessFlow
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

// GetN performs GET request for BusinessFlow collection, max N pages
func (r *ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequest) GetN(ctx context.Context, n int) ([]BusinessFlow, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BusinessFlow collection
func (r *ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequest) Get(ctx context.Context) ([]BusinessFlow, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BusinessFlow collection
func (r *ApprovalWorkflowProviderBusinessFlowsWithRequestsAwaitingMyDecisionCollectionRequest) Add(ctx context.Context, reqObj *BusinessFlow) (resObj *BusinessFlow, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PolicyTemplates returns request builder for GovernancePolicyTemplate collection
func (b *ApprovalWorkflowProviderRequestBuilder) PolicyTemplates() *ApprovalWorkflowProviderPolicyTemplatesCollectionRequestBuilder {
	bb := &ApprovalWorkflowProviderPolicyTemplatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policyTemplates"
	return bb
}

// ApprovalWorkflowProviderPolicyTemplatesCollectionRequestBuilder is request builder for GovernancePolicyTemplate collection
type ApprovalWorkflowProviderPolicyTemplatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernancePolicyTemplate collection
func (b *ApprovalWorkflowProviderPolicyTemplatesCollectionRequestBuilder) Request() *ApprovalWorkflowProviderPolicyTemplatesCollectionRequest {
	return &ApprovalWorkflowProviderPolicyTemplatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernancePolicyTemplate item
func (b *ApprovalWorkflowProviderPolicyTemplatesCollectionRequestBuilder) ID(id string) *GovernancePolicyTemplateRequestBuilder {
	bb := &GovernancePolicyTemplateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApprovalWorkflowProviderPolicyTemplatesCollectionRequest is request for GovernancePolicyTemplate collection
type ApprovalWorkflowProviderPolicyTemplatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernancePolicyTemplate collection
func (r *ApprovalWorkflowProviderPolicyTemplatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernancePolicyTemplate, error) {
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
	var values []GovernancePolicyTemplate
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
			value  []GovernancePolicyTemplate
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

// GetN performs GET request for GovernancePolicyTemplate collection, max N pages
func (r *ApprovalWorkflowProviderPolicyTemplatesCollectionRequest) GetN(ctx context.Context, n int) ([]GovernancePolicyTemplate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernancePolicyTemplate collection
func (r *ApprovalWorkflowProviderPolicyTemplatesCollectionRequest) Get(ctx context.Context) ([]GovernancePolicyTemplate, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernancePolicyTemplate collection
func (r *ApprovalWorkflowProviderPolicyTemplatesCollectionRequest) Add(ctx context.Context, reqObj *GovernancePolicyTemplate) (resObj *GovernancePolicyTemplate, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Requests returns request builder for RequestObject collection
func (b *ApprovalWorkflowProviderRequestBuilder) Requests() *ApprovalWorkflowProviderRequestsCollectionRequestBuilder {
	bb := &ApprovalWorkflowProviderRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/requests"
	return bb
}

// ApprovalWorkflowProviderRequestsCollectionRequestBuilder is request builder for RequestObject collection
type ApprovalWorkflowProviderRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RequestObject collection
func (b *ApprovalWorkflowProviderRequestsCollectionRequestBuilder) Request() *ApprovalWorkflowProviderRequestsCollectionRequest {
	return &ApprovalWorkflowProviderRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RequestObject item
func (b *ApprovalWorkflowProviderRequestsCollectionRequestBuilder) ID(id string) *RequestObjectRequestBuilder {
	bb := &RequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApprovalWorkflowProviderRequestsCollectionRequest is request for RequestObject collection
type ApprovalWorkflowProviderRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RequestObject collection
func (r *ApprovalWorkflowProviderRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RequestObject, error) {
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
	var values []RequestObject
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
			value  []RequestObject
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

// GetN performs GET request for RequestObject collection, max N pages
func (r *ApprovalWorkflowProviderRequestsCollectionRequest) GetN(ctx context.Context, n int) ([]RequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RequestObject collection
func (r *ApprovalWorkflowProviderRequestsCollectionRequest) Get(ctx context.Context) ([]RequestObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RequestObject collection
func (r *ApprovalWorkflowProviderRequestsCollectionRequest) Add(ctx context.Context, reqObj *RequestObject) (resObj *RequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RequestsAwaitingMyDecision returns request builder for RequestObject collection
func (b *ApprovalWorkflowProviderRequestBuilder) RequestsAwaitingMyDecision() *ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequestBuilder {
	bb := &ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/requestsAwaitingMyDecision"
	return bb
}

// ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequestBuilder is request builder for RequestObject collection
type ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RequestObject collection
func (b *ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequestBuilder) Request() *ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequest {
	return &ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RequestObject item
func (b *ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequestBuilder) ID(id string) *RequestObjectRequestBuilder {
	bb := &RequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequest is request for RequestObject collection
type ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RequestObject collection
func (r *ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RequestObject, error) {
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
	var values []RequestObject
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
			value  []RequestObject
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

// GetN performs GET request for RequestObject collection, max N pages
func (r *ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequest) GetN(ctx context.Context, n int) ([]RequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RequestObject collection
func (r *ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequest) Get(ctx context.Context) ([]RequestObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RequestObject collection
func (r *ApprovalWorkflowProviderRequestsAwaitingMyDecisionCollectionRequest) Add(ctx context.Context, reqObj *RequestObject) (resObj *RequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
