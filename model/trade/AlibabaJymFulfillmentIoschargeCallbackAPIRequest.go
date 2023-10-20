package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymFulfillmentIoschargeCallbackAPIRequest 代充充值回调 API请求
// alibaba.jym.fulfillment.ioscharge.callback
//
// 代充充值回调
type AlibabaJymFulfillmentIoschargeCallbackAPIRequest struct {
	model.Params
	// 充值回调请求
	_iosChargeCallbackRequestDto *IoschargeCallbackRequestDto
}

// NewAlibabaJymFulfillmentIoschargeCallbackRequest 初始化AlibabaJymFulfillmentIoschargeCallbackAPIRequest对象
func NewAlibabaJymFulfillmentIoschargeCallbackRequest() *AlibabaJymFulfillmentIoschargeCallbackAPIRequest {
	return &AlibabaJymFulfillmentIoschargeCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymFulfillmentIoschargeCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.fulfillment.ioscharge.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymFulfillmentIoschargeCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymFulfillmentIoschargeCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIosChargeCallbackRequestDto is IosChargeCallbackRequestDto Setter
// 充值回调请求
func (r *AlibabaJymFulfillmentIoschargeCallbackAPIRequest) SetIosChargeCallbackRequestDto(_iosChargeCallbackRequestDto *IoschargeCallbackRequestDto) error {
	r._iosChargeCallbackRequestDto = _iosChargeCallbackRequestDto
	r.Set("ios_charge_callback_request_dto", _iosChargeCallbackRequestDto)
	return nil
}

// GetIosChargeCallbackRequestDto IosChargeCallbackRequestDto Getter
func (r AlibabaJymFulfillmentIoschargeCallbackAPIRequest) GetIosChargeCallbackRequestDto() *IoschargeCallbackRequestDto {
	return r._iosChargeCallbackRequestDto
}
