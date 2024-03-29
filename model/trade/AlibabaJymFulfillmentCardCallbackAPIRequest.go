package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymFulfillmentCardCallbackAPIRequest 外部商家卡密结果回调 API请求
// alibaba.jym.fulfillment.card.callback
//
// 外部商家卡密结果回调
type AlibabaJymFulfillmentCardCallbackAPIRequest struct {
	model.Params
	// 卡密类结果回调请求
	_cardChargeCallbackRequestDto *CardChargeCallbackRequestDto
}

// NewAlibabaJymFulfillmentCardCallbackRequest 初始化AlibabaJymFulfillmentCardCallbackAPIRequest对象
func NewAlibabaJymFulfillmentCardCallbackRequest() *AlibabaJymFulfillmentCardCallbackAPIRequest {
	return &AlibabaJymFulfillmentCardCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymFulfillmentCardCallbackAPIRequest) Reset() {
	r._cardChargeCallbackRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymFulfillmentCardCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.fulfillment.card.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymFulfillmentCardCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymFulfillmentCardCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCardChargeCallbackRequestDto is CardChargeCallbackRequestDto Setter
// 卡密类结果回调请求
func (r *AlibabaJymFulfillmentCardCallbackAPIRequest) SetCardChargeCallbackRequestDto(_cardChargeCallbackRequestDto *CardChargeCallbackRequestDto) error {
	r._cardChargeCallbackRequestDto = _cardChargeCallbackRequestDto
	r.Set("card_charge_callback_request_dto", _cardChargeCallbackRequestDto)
	return nil
}

// GetCardChargeCallbackRequestDto CardChargeCallbackRequestDto Getter
func (r AlibabaJymFulfillmentCardCallbackAPIRequest) GetCardChargeCallbackRequestDto() *CardChargeCallbackRequestDto {
	return r._cardChargeCallbackRequestDto
}

var poolAlibabaJymFulfillmentCardCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymFulfillmentCardCallbackRequest()
	},
}

// GetAlibabaJymFulfillmentCardCallbackRequest 从 sync.Pool 获取 AlibabaJymFulfillmentCardCallbackAPIRequest
func GetAlibabaJymFulfillmentCardCallbackAPIRequest() *AlibabaJymFulfillmentCardCallbackAPIRequest {
	return poolAlibabaJymFulfillmentCardCallbackAPIRequest.Get().(*AlibabaJymFulfillmentCardCallbackAPIRequest)
}

// ReleaseAlibabaJymFulfillmentCardCallbackAPIRequest 将 AlibabaJymFulfillmentCardCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymFulfillmentCardCallbackAPIRequest(v *AlibabaJymFulfillmentCardCallbackAPIRequest) {
	v.Reset()
	poolAlibabaJymFulfillmentCardCallbackAPIRequest.Put(v)
}
