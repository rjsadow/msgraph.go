// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjsadow/msgraph.go/jsonx"
)

// AlertRequestBuilder is request builder for Alert
type AlertRequestBuilder struct{ BaseRequestBuilder }

// Request returns AlertRequest
func (b *AlertRequestBuilder) Request() *AlertRequest {
	return &AlertRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AlertRequest is request for Alert
type AlertRequest struct{ BaseRequest }

// Get performs GET request for Alert
func (r *AlertRequest) Get(ctx context.Context) (resObj *Alert, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Alert
func (r *AlertRequest) Update(ctx context.Context, reqObj *Alert) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Alert
func (r *AlertRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type AlertCollectionUpdateAlertsRequestBuilder struct{ BaseRequestBuilder }

// UpdateAlerts action undocumented
func (b *SecurityAlertsCollectionRequestBuilder) UpdateAlerts(reqObj *AlertCollectionUpdateAlertsRequestParameter) *AlertCollectionUpdateAlertsRequestBuilder {
	bb := &AlertCollectionUpdateAlertsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateAlerts"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AlertCollectionUpdateAlertsRequest struct{ BaseRequest }

//
func (b *AlertCollectionUpdateAlertsRequestBuilder) Request() *AlertCollectionUpdateAlertsRequest {
	return &AlertCollectionUpdateAlertsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AlertCollectionUpdateAlertsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Alert, error) {
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
	var values []Alert
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
			value  []Alert
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
func (r *AlertCollectionUpdateAlertsRequest) PostN(ctx context.Context, n int) ([]Alert, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *AlertCollectionUpdateAlertsRequest) Post(ctx context.Context) ([]Alert, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}
