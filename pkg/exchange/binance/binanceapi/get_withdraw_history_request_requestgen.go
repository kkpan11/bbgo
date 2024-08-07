// Code generated by "requestgen -method GET -url /sapi/v1/capital/withdraw/history -type GetWithdrawHistoryRequest -responseType []WithdrawRecord"; DO NOT EDIT.

package binanceapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func (g *GetWithdrawHistoryRequest) Coin(coin string) *GetWithdrawHistoryRequest {
	g.coin = coin
	return g
}

func (g *GetWithdrawHistoryRequest) WithdrawOrderId(withdrawOrderId string) *GetWithdrawHistoryRequest {
	g.withdrawOrderId = &withdrawOrderId
	return g
}

func (g *GetWithdrawHistoryRequest) Status(status WithdrawStatus) *GetWithdrawHistoryRequest {
	g.status = &status
	return g
}

func (g *GetWithdrawHistoryRequest) StartTime(startTime time.Time) *GetWithdrawHistoryRequest {
	g.startTime = &startTime
	return g
}

func (g *GetWithdrawHistoryRequest) EndTime(endTime time.Time) *GetWithdrawHistoryRequest {
	g.endTime = &endTime
	return g
}

func (g *GetWithdrawHistoryRequest) Limit(limit uint64) *GetWithdrawHistoryRequest {
	g.limit = &limit
	return g
}

func (g *GetWithdrawHistoryRequest) Offset(offset uint64) *GetWithdrawHistoryRequest {
	g.offset = &offset
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetWithdrawHistoryRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetWithdrawHistoryRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check coin field -> json key coin
	coin := g.coin

	// assign parameter of coin
	params["coin"] = coin
	// check withdrawOrderId field -> json key withdrawOrderId
	if g.withdrawOrderId != nil {
		withdrawOrderId := *g.withdrawOrderId

		// assign parameter of withdrawOrderId
		params["withdrawOrderId"] = withdrawOrderId
	} else {
	}
	// check status field -> json key status
	if g.status != nil {
		status := *g.status

		// TEMPLATE check-valid-values
		switch status {
		case WithdrawStatusEmailSent:
			params["status"] = status

		default:
			return nil, fmt.Errorf("status value %v is invalid", status)

		}
		// END TEMPLATE check-valid-values

		// assign parameter of status
		params["status"] = status
	} else {
	}
	// check startTime field -> json key startTime
	if g.startTime != nil {
		startTime := *g.startTime

		// assign parameter of startTime
		// convert time.Time to milliseconds time stamp
		params["startTime"] = strconv.FormatInt(startTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check endTime field -> json key endTime
	if g.endTime != nil {
		endTime := *g.endTime

		// assign parameter of endTime
		// convert time.Time to milliseconds time stamp
		params["endTime"] = strconv.FormatInt(endTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check limit field -> json key limit
	if g.limit != nil {
		limit := *g.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}
	// check offset field -> json key offset
	if g.offset != nil {
		offset := *g.offset

		// assign parameter of offset
		params["offset"] = offset
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetWithdrawHistoryRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetWithdrawHistoryRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetWithdrawHistoryRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetWithdrawHistoryRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetWithdrawHistoryRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetWithdrawHistoryRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetWithdrawHistoryRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

// GetPath returns the request path of the API
func (g *GetWithdrawHistoryRequest) GetPath() string {
	return "/sapi/v1/capital/withdraw/history"
}

// Do generates the request object and send the request object to the API endpoint
func (g *GetWithdrawHistoryRequest) Do(ctx context.Context) ([]WithdrawRecord, error) {

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	var apiURL string

	apiURL = g.GetPath()

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []WithdrawRecord

	type responseUnmarshaler interface {
		Unmarshal(data []byte) error
	}

	if unmarshaler, ok := interface{}(&apiResponse).(responseUnmarshaler); ok {
		if err := unmarshaler.Unmarshal(response.Body); err != nil {
			return nil, err
		}
	} else {
		// The line below checks the content type, however, some API server might not send the correct content type header,
		// Hence, this is commented for backward compatibility
		// response.IsJSON()
		if err := response.DecodeJSON(&apiResponse); err != nil {
			return nil, err
		}
	}

	type responseValidator interface {
		Validate() error
	}

	if validator, ok := interface{}(&apiResponse).(responseValidator); ok {
		if err := validator.Validate(); err != nil {
			return nil, err
		}
	}
	return apiResponse, nil
}
