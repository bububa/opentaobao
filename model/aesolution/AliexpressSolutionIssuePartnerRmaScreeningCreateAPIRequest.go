package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest aliexpress.solution.issue.partner.rma.screening.create API请求
// aliexpress.solution.issue.partner.rma.screening.create
//
// Receives information about screening results from after sales partners
type AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest struct {
	model.Params
	// Screening result creation request
	_screeningResultCreationRequest *RmaScreeningCreationRequest
}

// NewAliexpresssolutionissuepartnerrmascreeningcreateRequest 初始化AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest对象
func NewAliexpresssolutionissuepartnerrmascreeningcreateRequest() *AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest {
	return &AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.screening.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScreeningResultCreationRequest is ScreeningResultCreationRequest Setter
// Screening result creation request
func (r *AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest) SetScreeningResultCreationRequest(_screeningResultCreationRequest *RmaScreeningCreationRequest) error {
	r._screeningResultCreationRequest = _screeningResultCreationRequest
	r.Set("screening_result_creation_request", _screeningResultCreationRequest)
	return nil
}

// GetScreeningResultCreationRequest ScreeningResultCreationRequest Getter
func (r AliexpresssolutionissuepartnerrmascreeningcreateAPIRequest) GetScreeningResultCreationRequest() *RmaScreeningCreationRequest {
	return r._screeningResultCreationRequest
}
