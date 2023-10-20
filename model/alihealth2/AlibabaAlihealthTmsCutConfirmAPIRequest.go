package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtmscutconfirmAPIRequest 配拦截失败CP确认结果并回告 API请求
// alibaba.alihealth.tms.cut.confirm
//
// 配拦截失败CP确认结果并回告
type AlibabaalihealthtmscutconfirmAPIRequest struct {
	model.Params
	// 参数
	_tmsCutResultConfirmRequest *TmsCutResultConfirmRequest
}

// NewAlibabaalihealthtmscutconfirmRequest 初始化AlibabaalihealthtmscutconfirmAPIRequest对象
func NewAlibabaalihealthtmscutconfirmRequest() *AlibabaalihealthtmscutconfirmAPIRequest {
	return &AlibabaalihealthtmscutconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtmscutconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tms.cut.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtmscutconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtmscutconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsCutResultConfirmRequest is TmsCutResultConfirmRequest Setter
// 参数
func (r *AlibabaalihealthtmscutconfirmAPIRequest) SetTmsCutResultConfirmRequest(_tmsCutResultConfirmRequest *TmsCutResultConfirmRequest) error {
	r._tmsCutResultConfirmRequest = _tmsCutResultConfirmRequest
	r.Set("tms_cut_result_confirm_request", _tmsCutResultConfirmRequest)
	return nil
}

// GetTmsCutResultConfirmRequest TmsCutResultConfirmRequest Getter
func (r AlibabaalihealthtmscutconfirmAPIRequest) GetTmsCutResultConfirmRequest() *TmsCutResultConfirmRequest {
	return r._tmsCutResultConfirmRequest
}
