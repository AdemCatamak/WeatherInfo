package weatherServices

import "errors"

type Error struct {
	error
}

func NewError(errMsg string) Error {
	return Error{error: errors.New(errMsg)}
}

var (
	ErrorLocationNotFound           Error = NewError("Location Not Found")
	ErrorUnauthorizedRequest        Error = NewError("Unauthorized Request")
	ErrorWeatherServiceGatewayError Error = NewError("Weather Service Gateway Error")
)
