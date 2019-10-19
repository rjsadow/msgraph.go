// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UserAnalyticsRequestBuilder is request builder for UserAnalytics
type UserAnalyticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserAnalyticsRequest
func (b *UserAnalyticsRequestBuilder) Request() *UserAnalyticsRequest {
	return &UserAnalyticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserAnalyticsRequest is request for UserAnalytics
type UserAnalyticsRequest struct{ BaseRequest }

// Do performs HTTP request for UserAnalytics
func (r *UserAnalyticsRequest) Do(method, path string, reqObj interface{}) (resObj *UserAnalytics, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for UserAnalytics
func (r *UserAnalyticsRequest) Get() (*UserAnalytics, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for UserAnalytics
func (r *UserAnalyticsRequest) Update(reqObj *UserAnalytics) (*UserAnalytics, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for UserAnalytics
func (r *UserAnalyticsRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ActivityStatistics returns request builder for ActivityStatistics collection
func (b *UserAnalyticsRequestBuilder) ActivityStatistics() *UserAnalyticsActivityStatisticsCollectionRequestBuilder {
	bb := &UserAnalyticsActivityStatisticsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activityStatistics"
	return bb
}

// UserAnalyticsActivityStatisticsCollectionRequestBuilder is request builder for ActivityStatistics collection
type UserAnalyticsActivityStatisticsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ActivityStatistics collection
func (b *UserAnalyticsActivityStatisticsCollectionRequestBuilder) Request() *UserAnalyticsActivityStatisticsCollectionRequest {
	return &UserAnalyticsActivityStatisticsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ActivityStatistics item
func (b *UserAnalyticsActivityStatisticsCollectionRequestBuilder) ID(id string) *ActivityStatisticsRequestBuilder {
	bb := &ActivityStatisticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UserAnalyticsActivityStatisticsCollectionRequest is request for ActivityStatistics collection
type UserAnalyticsActivityStatisticsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ActivityStatistics, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionRequest) Paging(method, path string, obj interface{}) ([]ActivityStatistics, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ActivityStatistics
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ActivityStatistics
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionRequest) Get() ([]ActivityStatistics, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionRequest) Add(reqObj *ActivityStatistics) (*ActivityStatistics, error) {
	return r.Do("POST", "", reqObj)
}