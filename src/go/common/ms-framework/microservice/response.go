package microservice

import (
	"com.gft.tsbo-training.src.go/common/ms-framework/dispatcher"
)

// ###########################################################################
// ###########################################################################
// MicroService Response
// ###########################################################################
// ###########################################################################

// Response encapsulates the data of a general ms response
type Response struct {
	dispatcher.Response
}

// ###########################################################################

// InitResponseFromMicroService ...
func InitResponseFromMicroService(r *Response, ms IConfiguration, status string) {
	dispatcher.InitResponseFromDispatcher(&r.Response, ms, status)
}
