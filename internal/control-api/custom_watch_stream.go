package control_api

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var Logger *zap.SugaredLogger = nil

func (r ApiWatchEventsRequest) WatchEventsStream() (*WatchEventsStream, *http.Response, error) {
	return r.ApiService.WatchEventsStream(r)
}

func (a *VPCApiService) WatchEventsStream(r ApiWatchEventsRequest) (*WatchEventsStream, *http.Response, error) {
	// this is a stream friendly version of WatchEventsExecute
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WatchEventsStream // different from WatchEventsExecute
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VPCApiService.WatchEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/vpcs/{id}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.watches == nil {
		return localVarReturnValue, nil, reportError("watches is required and must be specified")
	}

	if r.publicKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "public_key", r.publicKey, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.watches
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {

		// different from WatchEventsExecute: start
		localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
		_ = localVarHTTPResponse.Body.Close()
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
		if err != nil {
			return nil, localVarHTTPResponse, err
		}
		// different from WatchEventsExecute: end

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelsBaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelsBaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelsBaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelsBaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	// different from WatchEventsExecute: start
	return &WatchEventsStream{
		close:   localVarHTTPResponse.Body.Close,
		decoder: json.NewDecoder(localVarHTTPResponse.Body),
	}, localVarHTTPResponse, nil
	// different from WatchEventsExecute: end

}

type WatchEventsStream struct {
	decoder *json.Decoder
	close   func() error
}

func (ds *WatchEventsStream) Receive() (ModelsWatchEvent, error) {
	event := ModelsWatchEvent{}
	err := ds.decoder.Decode(&event)
	return event, err
}

func (ds *WatchEventsStream) Close() error {
	return ds.close()
}
