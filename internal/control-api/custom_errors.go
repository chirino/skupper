package control_api

import (
	"errors"
	"fmt"
	"net/http"
)

func Simplify[T any](resp T, httpResp *http.Response, err error) (T, error) {
	if err != nil {
		status := httpResp.StatusCode
		var openAPIError *GenericOpenAPIError
		switch {
		case errors.As(err, &openAPIError):
			model := openAPIError.Model()
			switch err := model.(type) {
			case ModelsBaseError:
				return resp, fmt.Errorf("error: %s, status: %d", err.GetError(), status)
			case ModelsConflictsError:
				return resp, fmt.Errorf("error: %s: conflicting id: %s, status: %d", err.GetError(), err.GetId(), status)
			case ModelsNotAllowedError:
				message := fmt.Sprintf("error: %s", err.GetError())
				if err.Reason != nil {
					message += fmt.Sprintf(", reason: %s", err.GetReason())
				}
				message += fmt.Sprintf(", status: %d", status)
				return resp, fmt.Errorf(message)
			case ModelsValidationError:
				message := fmt.Sprintf("error: %s", err.GetError())
				if err.Field != nil {
					message += fmt.Sprintf(", field: %s", err.GetField())
				}
				message += fmt.Sprintf(", status: %d", status)
				return resp, fmt.Errorf(message)
			case ModelsInternalServerError:
				return resp, fmt.Errorf("error: %s: trace id: %s, status: %d", err.GetError(), err.GetTraceId(), status)
			default:
				return resp, fmt.Errorf("error: %s, status: %d", string(openAPIError.Body()), status)
			}
		default:
			return resp, fmt.Errorf("error: %+v, status: %d", err, status)
		}
		errors.Join()
	}
	return resp, nil
}
