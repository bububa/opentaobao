package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymfulfillmentcardcallbackAPIRequest 外部商家卡密结果回调 API请求
// alibaba.jym.fulfillment.card.callback
//
// 外部商家卡密结果回调
type AlibabajymfulfillmentcardcallbackAPIRequest struct {
	model.Params
	// 卡密类结果回调请求
	_cardChargeCallbackRequestDto *CardChargeCallbackRequestDto
}

// NewAlibabajymfulfillmentcardcallbackRequest 初始化AlibabajymfulfillmentcardcallbackAPIRequest对象
func NewAlibabajymfulfillmentcardcallbackRequest() *AlibabajymfulfillmentcardcallbackAPIRequest {
	return &AlibabajymfulfillmentcardcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymfulfillmentcardcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.fulfillment.card.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymfulfillmentcardcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymfulfillmentcardcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCardChargeCallbackRequestDto is CardChargeCallbackRequestDto Setter
// 卡密类结果回调请求
func (r *AlibabajymfulfillmentcardcallbackAPIRequest) SetCardChargeCallbackRequestDto(_cardChargeCallbackRequestDto *CardChargeCallbackRequestDto) error {
	r._cardChargeCallbackRequestDto = _cardChargeCallbackRequestDto
	r.Set("card_charge_callback_request_dto", _cardChargeCallbackRequestDto)
	return nil
}

// GetCardChargeCallbackRequestDto CardChargeCallbackRequestDto Getter
func (r AlibabajymfulfillmentcardcallbackAPIRequest) GetCardChargeCallbackRequestDto() *CardChargeCallbackRequestDto {
	return r._cardChargeCallbackRequestDto
}
