// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjsadow/msgraph.go/jsonx"
)

// OutlookCategoryRequestBuilder is request builder for OutlookCategory
type OutlookCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookCategoryRequest
func (b *OutlookCategoryRequestBuilder) Request() *OutlookCategoryRequest {
	return &OutlookCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookCategoryRequest is request for OutlookCategory
type OutlookCategoryRequest struct{ BaseRequest }

// Get performs GET request for OutlookCategory
func (r *OutlookCategoryRequest) Get(ctx context.Context) (resObj *OutlookCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OutlookCategory
func (r *OutlookCategoryRequest) Update(ctx context.Context, reqObj *OutlookCategory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OutlookCategory
func (r *OutlookCategoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OutlookItemRequestBuilder is request builder for OutlookItem
type OutlookItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookItemRequest
func (b *OutlookItemRequestBuilder) Request() *OutlookItemRequest {
	return &OutlookItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookItemRequest is request for OutlookItem
type OutlookItemRequest struct{ BaseRequest }

// Get performs GET request for OutlookItem
func (r *OutlookItemRequest) Get(ctx context.Context) (resObj *OutlookItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OutlookItem
func (r *OutlookItemRequest) Update(ctx context.Context, reqObj *OutlookItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OutlookItem
func (r *OutlookItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OutlookTaskRequestBuilder is request builder for OutlookTask
type OutlookTaskRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookTaskRequest
func (b *OutlookTaskRequestBuilder) Request() *OutlookTaskRequest {
	return &OutlookTaskRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookTaskRequest is request for OutlookTask
type OutlookTaskRequest struct{ BaseRequest }

// Get performs GET request for OutlookTask
func (r *OutlookTaskRequest) Get(ctx context.Context) (resObj *OutlookTask, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OutlookTask
func (r *OutlookTaskRequest) Update(ctx context.Context, reqObj *OutlookTask) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OutlookTask
func (r *OutlookTaskRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OutlookTaskFolderRequestBuilder is request builder for OutlookTaskFolder
type OutlookTaskFolderRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookTaskFolderRequest
func (b *OutlookTaskFolderRequestBuilder) Request() *OutlookTaskFolderRequest {
	return &OutlookTaskFolderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookTaskFolderRequest is request for OutlookTaskFolder
type OutlookTaskFolderRequest struct{ BaseRequest }

// Get performs GET request for OutlookTaskFolder
func (r *OutlookTaskFolderRequest) Get(ctx context.Context) (resObj *OutlookTaskFolder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OutlookTaskFolder
func (r *OutlookTaskFolderRequest) Update(ctx context.Context, reqObj *OutlookTaskFolder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OutlookTaskFolder
func (r *OutlookTaskFolderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OutlookTaskGroupRequestBuilder is request builder for OutlookTaskGroup
type OutlookTaskGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookTaskGroupRequest
func (b *OutlookTaskGroupRequestBuilder) Request() *OutlookTaskGroupRequest {
	return &OutlookTaskGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookTaskGroupRequest is request for OutlookTaskGroup
type OutlookTaskGroupRequest struct{ BaseRequest }

// Get performs GET request for OutlookTaskGroup
func (r *OutlookTaskGroupRequest) Get(ctx context.Context) (resObj *OutlookTaskGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OutlookTaskGroup
func (r *OutlookTaskGroupRequest) Update(ctx context.Context, reqObj *OutlookTaskGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OutlookTaskGroup
func (r *OutlookTaskGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OutlookUserRequestBuilder is request builder for OutlookUser
type OutlookUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookUserRequest
func (b *OutlookUserRequestBuilder) Request() *OutlookUserRequest {
	return &OutlookUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookUserRequest is request for OutlookUser
type OutlookUserRequest struct{ BaseRequest }

// Get performs GET request for OutlookUser
func (r *OutlookUserRequest) Get(ctx context.Context) (resObj *OutlookUser, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OutlookUser
func (r *OutlookUserRequest) Update(ctx context.Context, reqObj *OutlookUser) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OutlookUser
func (r *OutlookUserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type OutlookTaskCompleteRequestBuilder struct{ BaseRequestBuilder }

// Complete action undocumented
func (b *OutlookTaskRequestBuilder) Complete(reqObj *OutlookTaskCompleteRequestParameter) *OutlookTaskCompleteRequestBuilder {
	bb := &OutlookTaskCompleteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/complete"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OutlookTaskCompleteRequest struct{ BaseRequest }

//
func (b *OutlookTaskCompleteRequestBuilder) Request() *OutlookTaskCompleteRequest {
	return &OutlookTaskCompleteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OutlookTaskCompleteRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OutlookTask, error) {
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
	var values []OutlookTask
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
			value  []OutlookTask
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
func (r *OutlookTaskCompleteRequest) PostN(ctx context.Context, n int) ([]OutlookTask, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *OutlookTaskCompleteRequest) Post(ctx context.Context) ([]OutlookTask, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}
