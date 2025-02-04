// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjsadow/msgraph.go/jsonx"
)

// Controls returns request builder for ProgramControl collection
func (b *ProgramRequestBuilder) Controls() *ProgramControlsCollectionRequestBuilder {
	bb := &ProgramControlsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/controls"
	return bb
}

// ProgramControlsCollectionRequestBuilder is request builder for ProgramControl collection
type ProgramControlsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ProgramControl collection
func (b *ProgramControlsCollectionRequestBuilder) Request() *ProgramControlsCollectionRequest {
	return &ProgramControlsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ProgramControl item
func (b *ProgramControlsCollectionRequestBuilder) ID(id string) *ProgramControlRequestBuilder {
	bb := &ProgramControlRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ProgramControlsCollectionRequest is request for ProgramControl collection
type ProgramControlsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ProgramControl, error) {
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
	var values []ProgramControl
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
			value  []ProgramControl
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

// GetN performs GET request for ProgramControl collection, max N pages
func (r *ProgramControlsCollectionRequest) GetN(ctx context.Context, n int) ([]ProgramControl, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Get(ctx context.Context) ([]ProgramControl, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Add(ctx context.Context, reqObj *ProgramControl) (resObj *ProgramControl, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Program is navigation property
func (b *ProgramControlRequestBuilder) Program() *ProgramRequestBuilder {
	bb := &ProgramRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/program"
	return bb
}
