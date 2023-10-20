package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymfulfillmentioschargecallbackAPIRequest 代充充值回调 API请求
// alibaba.jym.fulfillment.ioscharge.callback
//
// 代充充值回调
type AlibabajymfulfillmentioschargecallbackAPIRequest struct {
	model.Params
	// 充值回调请求
	_iosChargeCallbackRequestDto *IoschargeCallbackRequestDto
}

// NewAlibabajymfulfillmentioschargecallbackRequest 初始化AlibabajymfulfillmentioschargecallbackAPIRequest对象
func NewAlibabajymfulfillmentioschargecallbackRequest() *AlibabajymfulfillmentioschargecallbackAPIRequest {
	return &AlibabajymfulfillmentioschargecallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymfulfillmentioschargecallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.fulfillment.ioscharge.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymfulfillmentioschargecallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymfulfillmentioschargecallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIosChargeCallbackRequestDto is IosChargeCallbackRequestDto Setter
// 充值回调请求
func (r *AlibabajymfulfillmentioschargecallbackAPIRequest) SetIosChargeCallbackRequestDto(_iosChargeCallbackRequestDto *IoschargeCallbackRequestDto) error {
	r._iosChargeCallbackRequestDto = _iosChargeCallbackRequestDto
	r.Set("ios_charge_callback_request_dto", _iosChargeCallbackRequestDto)
	return nil
}

// GetIosChargeCallbackRequestDto IosChargeCallbackRequestDto Getter
func (r AlibabajymfulfillmentioschargecallbackAPIRequest) GetIosChargeCallbackRequestDto() *IoschargeCallbackRequestDto {
	return r._iosChargeCallbackRequestDto
}
