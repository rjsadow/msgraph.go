// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjsadow/msgraph.go/jsonx"
)

// SynchronizationJobCollectionValidateCredentialsRequestParameter undocumented
type SynchronizationJobCollectionValidateCredentialsRequestParameter struct {
	// ApplicationIdentifier undocumented
	ApplicationIdentifier *string `json:"applicationIdentifier,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// UseSavedCredentials undocumented
	UseSavedCredentials *bool `json:"useSavedCredentials,omitempty"`
	// Credentials undocumented
	Credentials []SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
}

// SynchronizationJobPauseRequestParameter undocumented
type SynchronizationJobPauseRequestParameter struct {
}

// SynchronizationJobStartRequestParameter undocumented
type SynchronizationJobStartRequestParameter struct {
}

// SynchronizationJobStopRequestParameter undocumented
type SynchronizationJobStopRequestParameter struct {
}

// SynchronizationJobApplyRequestParameter undocumented
type SynchronizationJobApplyRequestParameter struct {
	// ObjectID undocumented
	ObjectID *string `json:"objectId,omitempty"`
	// TypeName undocumented
	TypeName *string `json:"typeName,omitempty"`
	// RuleID undocumented
	RuleID *string `json:"ruleId,omitempty"`
}

// SynchronizationJobRestartRequestParameter undocumented
type SynchronizationJobRestartRequestParameter struct {
	// Criteria undocumented
	Criteria *SynchronizationJobRestartCriteria `json:"criteria,omitempty"`
}

// SynchronizationJobValidateCredentialsRequestParameter undocumented
type SynchronizationJobValidateCredentialsRequestParameter struct {
	// ApplicationIdentifier undocumented
	ApplicationIdentifier *string `json:"applicationIdentifier,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// UseSavedCredentials undocumented
	UseSavedCredentials *bool `json:"useSavedCredentials,omitempty"`
	// Credentials undocumented
	Credentials []SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
}

// SynchronizationSchemaParseExpressionRequestParameter undocumented
type SynchronizationSchemaParseExpressionRequestParameter struct {
	// Expression undocumented
	Expression *string `json:"expression,omitempty"`
	// TestInputObject undocumented
	TestInputObject *ExpressionInputObject `json:"testInputObject,omitempty"`
	// TargetAttributeDefinition undocumented
	TargetAttributeDefinition *AttributeDefinition `json:"targetAttributeDefinition,omitempty"`
}

// Jobs returns request builder for SynchronizationJob collection
func (b *SynchronizationRequestBuilder) Jobs() *SynchronizationJobsCollectionRequestBuilder {
	bb := &SynchronizationJobsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/jobs"
	return bb
}

// SynchronizationJobsCollectionRequestBuilder is request builder for SynchronizationJob collection
type SynchronizationJobsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SynchronizationJob collection
func (b *SynchronizationJobsCollectionRequestBuilder) Request() *SynchronizationJobsCollectionRequest {
	return &SynchronizationJobsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SynchronizationJob item
func (b *SynchronizationJobsCollectionRequestBuilder) ID(id string) *SynchronizationJobRequestBuilder {
	bb := &SynchronizationJobRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SynchronizationJobsCollectionRequest is request for SynchronizationJob collection
type SynchronizationJobsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SynchronizationJob collection
func (r *SynchronizationJobsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SynchronizationJob, error) {
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
	var values []SynchronizationJob
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
			value  []SynchronizationJob
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

// GetN performs GET request for SynchronizationJob collection, max N pages
func (r *SynchronizationJobsCollectionRequest) GetN(ctx context.Context, n int) ([]SynchronizationJob, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SynchronizationJob collection
func (r *SynchronizationJobsCollectionRequest) Get(ctx context.Context) ([]SynchronizationJob, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SynchronizationJob collection
func (r *SynchronizationJobsCollectionRequest) Add(ctx context.Context, reqObj *SynchronizationJob) (resObj *SynchronizationJob, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Templates returns request builder for SynchronizationTemplate collection
func (b *SynchronizationRequestBuilder) Templates() *SynchronizationTemplatesCollectionRequestBuilder {
	bb := &SynchronizationTemplatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/templates"
	return bb
}

// SynchronizationTemplatesCollectionRequestBuilder is request builder for SynchronizationTemplate collection
type SynchronizationTemplatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SynchronizationTemplate collection
func (b *SynchronizationTemplatesCollectionRequestBuilder) Request() *SynchronizationTemplatesCollectionRequest {
	return &SynchronizationTemplatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SynchronizationTemplate item
func (b *SynchronizationTemplatesCollectionRequestBuilder) ID(id string) *SynchronizationTemplateRequestBuilder {
	bb := &SynchronizationTemplateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SynchronizationTemplatesCollectionRequest is request for SynchronizationTemplate collection
type SynchronizationTemplatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SynchronizationTemplate collection
func (r *SynchronizationTemplatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SynchronizationTemplate, error) {
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
	var values []SynchronizationTemplate
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
			value  []SynchronizationTemplate
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

// GetN performs GET request for SynchronizationTemplate collection, max N pages
func (r *SynchronizationTemplatesCollectionRequest) GetN(ctx context.Context, n int) ([]SynchronizationTemplate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SynchronizationTemplate collection
func (r *SynchronizationTemplatesCollectionRequest) Get(ctx context.Context) ([]SynchronizationTemplate, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SynchronizationTemplate collection
func (r *SynchronizationTemplatesCollectionRequest) Add(ctx context.Context, reqObj *SynchronizationTemplate) (resObj *SynchronizationTemplate, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Schema is navigation property
func (b *SynchronizationJobRequestBuilder) Schema() *SynchronizationSchemaRequestBuilder {
	bb := &SynchronizationSchemaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schema"
	return bb
}

// Directories returns request builder for DirectoryDefinition collection
func (b *SynchronizationSchemaRequestBuilder) Directories() *SynchronizationSchemaDirectoriesCollectionRequestBuilder {
	bb := &SynchronizationSchemaDirectoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directories"
	return bb
}

// SynchronizationSchemaDirectoriesCollectionRequestBuilder is request builder for DirectoryDefinition collection
type SynchronizationSchemaDirectoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryDefinition collection
func (b *SynchronizationSchemaDirectoriesCollectionRequestBuilder) Request() *SynchronizationSchemaDirectoriesCollectionRequest {
	return &SynchronizationSchemaDirectoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryDefinition item
func (b *SynchronizationSchemaDirectoriesCollectionRequestBuilder) ID(id string) *DirectoryDefinitionRequestBuilder {
	bb := &DirectoryDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SynchronizationSchemaDirectoriesCollectionRequest is request for DirectoryDefinition collection
type SynchronizationSchemaDirectoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryDefinition collection
func (r *SynchronizationSchemaDirectoriesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryDefinition, error) {
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
	var values []DirectoryDefinition
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
			value  []DirectoryDefinition
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

// GetN performs GET request for DirectoryDefinition collection, max N pages
func (r *SynchronizationSchemaDirectoriesCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryDefinition collection
func (r *SynchronizationSchemaDirectoriesCollectionRequest) Get(ctx context.Context) ([]DirectoryDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryDefinition collection
func (r *SynchronizationSchemaDirectoriesCollectionRequest) Add(ctx context.Context, reqObj *DirectoryDefinition) (resObj *DirectoryDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Schema is navigation property
func (b *SynchronizationTemplateRequestBuilder) Schema() *SynchronizationSchemaRequestBuilder {
	bb := &SynchronizationSchemaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schema"
	return bb
}
