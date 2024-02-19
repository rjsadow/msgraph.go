// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjsadow/msgraph.go/jsonx"
)

// InformationProtectionRequestBuilder is request builder for InformationProtection
type InformationProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns InformationProtectionRequest
func (b *InformationProtectionRequestBuilder) Request() *InformationProtectionRequest {
	return &InformationProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InformationProtectionRequest is request for InformationProtection
type InformationProtectionRequest struct{ BaseRequest }

// Get performs GET request for InformationProtection
func (r *InformationProtectionRequest) Get(ctx context.Context) (resObj *InformationProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InformationProtection
func (r *InformationProtectionRequest) Update(ctx context.Context, reqObj *InformationProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InformationProtection
func (r *InformationProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InformationProtectionLabelRequestBuilder is request builder for InformationProtectionLabel
type InformationProtectionLabelRequestBuilder struct{ BaseRequestBuilder }

// Request returns InformationProtectionLabelRequest
func (b *InformationProtectionLabelRequestBuilder) Request() *InformationProtectionLabelRequest {
	return &InformationProtectionLabelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InformationProtectionLabelRequest is request for InformationProtectionLabel
type InformationProtectionLabelRequest struct{ BaseRequest }

// Get performs GET request for InformationProtectionLabel
func (r *InformationProtectionLabelRequest) Get(ctx context.Context) (resObj *InformationProtectionLabel, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InformationProtectionLabel
func (r *InformationProtectionLabelRequest) Update(ctx context.Context, reqObj *InformationProtectionLabel) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InformationProtectionLabel
func (r *InformationProtectionLabelRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InformationProtectionPolicyRequestBuilder is request builder for InformationProtectionPolicy
type InformationProtectionPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns InformationProtectionPolicyRequest
func (b *InformationProtectionPolicyRequestBuilder) Request() *InformationProtectionPolicyRequest {
	return &InformationProtectionPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InformationProtectionPolicyRequest is request for InformationProtectionPolicy
type InformationProtectionPolicyRequest struct{ BaseRequest }

// Get performs GET request for InformationProtectionPolicy
func (r *InformationProtectionPolicyRequest) Get(ctx context.Context) (resObj *InformationProtectionPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InformationProtectionPolicy
func (r *InformationProtectionPolicyRequest) Update(ctx context.Context, reqObj *InformationProtectionPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InformationProtectionPolicy
func (r *InformationProtectionPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type InformationProtectionLabelCollectionExtractLabelRequestBuilder struct{ BaseRequestBuilder }

// ExtractLabel action undocumented
func (b *InformationProtectionPolicyLabelsCollectionRequestBuilder) ExtractLabel(reqObj *InformationProtectionLabelCollectionExtractLabelRequestParameter) *InformationProtectionLabelCollectionExtractLabelRequestBuilder {
	bb := &InformationProtectionLabelCollectionExtractLabelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/extractLabel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionLabelCollectionExtractLabelRequest struct{ BaseRequest }

//
func (b *InformationProtectionLabelCollectionExtractLabelRequestBuilder) Request() *InformationProtectionLabelCollectionExtractLabelRequest {
	return &InformationProtectionLabelCollectionExtractLabelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionLabelCollectionExtractLabelRequest) Post(ctx context.Context) (resObj *InformationProtectionContentLabel, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type InformationProtectionLabelCollectionEvaluateApplicationRequestBuilder struct{ BaseRequestBuilder }

// EvaluateApplication action undocumented
func (b *InformationProtectionPolicyLabelsCollectionRequestBuilder) EvaluateApplication(reqObj *InformationProtectionLabelCollectionEvaluateApplicationRequestParameter) *InformationProtectionLabelCollectionEvaluateApplicationRequestBuilder {
	bb := &InformationProtectionLabelCollectionEvaluateApplicationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateApplication"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionLabelCollectionEvaluateApplicationRequest struct{ BaseRequest }

//
func (b *InformationProtectionLabelCollectionEvaluateApplicationRequestBuilder) Request() *InformationProtectionLabelCollectionEvaluateApplicationRequest {
	return &InformationProtectionLabelCollectionEvaluateApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionLabelCollectionEvaluateApplicationRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]InformationProtectionAction, error) {
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
	var values []InformationProtectionAction
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
			value  []InformationProtectionAction
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
func (r *InformationProtectionLabelCollectionEvaluateApplicationRequest) PostN(ctx context.Context, n int) ([]InformationProtectionAction, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *InformationProtectionLabelCollectionEvaluateApplicationRequest) Post(ctx context.Context) ([]InformationProtectionAction, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type InformationProtectionLabelCollectionEvaluateRemovalRequestBuilder struct{ BaseRequestBuilder }

// EvaluateRemoval action undocumented
func (b *InformationProtectionPolicyLabelsCollectionRequestBuilder) EvaluateRemoval(reqObj *InformationProtectionLabelCollectionEvaluateRemovalRequestParameter) *InformationProtectionLabelCollectionEvaluateRemovalRequestBuilder {
	bb := &InformationProtectionLabelCollectionEvaluateRemovalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateRemoval"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionLabelCollectionEvaluateRemovalRequest struct{ BaseRequest }

//
func (b *InformationProtectionLabelCollectionEvaluateRemovalRequestBuilder) Request() *InformationProtectionLabelCollectionEvaluateRemovalRequest {
	return &InformationProtectionLabelCollectionEvaluateRemovalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionLabelCollectionEvaluateRemovalRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]InformationProtectionAction, error) {
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
	var values []InformationProtectionAction
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
			value  []InformationProtectionAction
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
func (r *InformationProtectionLabelCollectionEvaluateRemovalRequest) PostN(ctx context.Context, n int) ([]InformationProtectionAction, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *InformationProtectionLabelCollectionEvaluateRemovalRequest) Post(ctx context.Context) ([]InformationProtectionAction, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type InformationProtectionLabelCollectionEvaluateClassificationResultsRequestBuilder struct{ BaseRequestBuilder }

// EvaluateClassificationResults action undocumented
func (b *InformationProtectionPolicyLabelsCollectionRequestBuilder) EvaluateClassificationResults(reqObj *InformationProtectionLabelCollectionEvaluateClassificationResultsRequestParameter) *InformationProtectionLabelCollectionEvaluateClassificationResultsRequestBuilder {
	bb := &InformationProtectionLabelCollectionEvaluateClassificationResultsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateClassificationResults"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionLabelCollectionEvaluateClassificationResultsRequest struct{ BaseRequest }

//
func (b *InformationProtectionLabelCollectionEvaluateClassificationResultsRequestBuilder) Request() *InformationProtectionLabelCollectionEvaluateClassificationResultsRequest {
	return &InformationProtectionLabelCollectionEvaluateClassificationResultsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionLabelCollectionEvaluateClassificationResultsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]InformationProtectionAction, error) {
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
	var values []InformationProtectionAction
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
			value  []InformationProtectionAction
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
func (r *InformationProtectionLabelCollectionEvaluateClassificationResultsRequest) PostN(ctx context.Context, n int) ([]InformationProtectionAction, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *InformationProtectionLabelCollectionEvaluateClassificationResultsRequest) Post(ctx context.Context) ([]InformationProtectionAction, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type InformationProtectionEvaluateLabelsAndPoliciesRequestBuilder struct{ BaseRequestBuilder }

// EvaluateLabelsAndPolicies action undocumented
func (b *InformationProtectionRequestBuilder) EvaluateLabelsAndPolicies(reqObj *InformationProtectionEvaluateLabelsAndPoliciesRequestParameter) *InformationProtectionEvaluateLabelsAndPoliciesRequestBuilder {
	bb := &InformationProtectionEvaluateLabelsAndPoliciesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateLabelsAndPolicies"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionEvaluateLabelsAndPoliciesRequest struct{ BaseRequest }

//
func (b *InformationProtectionEvaluateLabelsAndPoliciesRequestBuilder) Request() *InformationProtectionEvaluateLabelsAndPoliciesRequest {
	return &InformationProtectionEvaluateLabelsAndPoliciesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionEvaluateLabelsAndPoliciesRequest) Post(ctx context.Context) (resObj *EvaluateLabelsAndPoliciesJobResponse, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
