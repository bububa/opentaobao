package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymFulfillmentIoschargeCallbackAPIRequest 代充充值回调 API请求
// alibaba.jym.fulfillment.ioscharge.callback
//
// 代充充值回调
type AlibabaJymFulfillmentIoschargeCallbackAPIRequest struct {
	model.Params
	// 充值回调请求
	_iosChargeCallbackRequestDto *IOSChargeCallbackRequestDto
}

// NewAlibabaJymFulfillmentIoschargeCallbackRequest 初始化AlibabaJymFulfillmentIoschargeCallbackAPIRequest对象
func NewAlibabaJymFulfillmentIoschargeCallbackRequest() *AlibabaJymFulfillmentIoschargeCallbackAPIRequest {
	return &AlibabaJymFulfillmentIoschargeCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymFulfillmentIoschargeCallbackAPIRequest) Reset() {
	r._iosChargeCallbackRequestDto = nil
	r.Params.ToZero()
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
func (r *AlibabaJymFulfillmentIoschargeCallbackAPIRequest) SetIosChargeCallbackRequestDto(_iosChargeCallbackRequestDto *IOSChargeCallbackRequestDto) error {
	r._iosChargeCallbackRequestDto = _iosChargeCallbackRequestDto
	r.Set("ios_charge_callback_request_dto", _iosChargeCallbackRequestDto)
	return nil
}

// GetIosChargeCallbackRequestDto IosChargeCallbackRequestDto Getter
func (r AlibabaJymFulfillmentIoschargeCallbackAPIRequest) GetIosChargeCallbackRequestDto() *IOSChargeCallbackRequestDto {
	return r._iosChargeCallbackRequestDto
}

var poolAlibabaJymFulfillmentIoschargeCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymFulfillmentIoschargeCallbackRequest()
	},
}

// GetAlibabaJymFulfillmentIoschargeCallbackRequest 从 sync.Pool 获取 AlibabaJymFulfillmentIoschargeCallbackAPIRequest
func GetAlibabaJymFulfillmentIoschargeCallbackAPIRequest() *AlibabaJymFulfillmentIoschargeCallbackAPIRequest {
	return poolAlibabaJymFulfillmentIoschargeCallbackAPIRequest.Get().(*AlibabaJymFulfillmentIoschargeCallbackAPIRequest)
}

// ReleaseAlibabaJymFulfillmentIoschargeCallbackAPIRequest 将 AlibabaJymFulfillmentIoschargeCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymFulfillmentIoschargeCallbackAPIRequest(v *AlibabaJymFulfillmentIoschargeCallbackAPIRequest) {
	v.Reset()
	poolAlibabaJymFulfillmentIoschargeCallbackAPIRequest.Put(v)
}
