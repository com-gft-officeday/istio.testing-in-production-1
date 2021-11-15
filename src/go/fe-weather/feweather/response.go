package fe-weather

import (
	"com.gft.tsbo-training.src.go/common/ms-framework/microservice"
)

// ###########################################################################
// ###########################################################################
// FeMeasure Response - Measure
// ###########################################################################
// ###########################################################################

// MeasureResponse Encapsulates the reploy of fe-weather
type MeasureResponse struct {
	microservice.Response
	RnrSvcVersion string `json:"rnrSvcVersion"`
}

// ###########################################################################

// InitMeasureResponse Constructor of a response of fe-weather
func InitMeasureResponse(r *MeasureResponse, status string, ms *FeMeasure) {
	microservice.InitResponseFromMicroService(&r.Response, ms, status)
	r.RnrSvcVersion = "???"
}

// NewMeasureResponse ...
func NewMeasureResponse(status string, ms *FeMeasure) *MeasureResponse {
	var r MeasureResponse
	InitMeasureResponse(&r, status, ms)
	return &r
}
