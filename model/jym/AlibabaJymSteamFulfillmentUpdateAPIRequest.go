package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymSteamFulfillmentUpdateAPIRequest 交易猫Steam类目发履约态变更 API请求
// alibaba.jym.steam.fulfillment.update
//
// 交易猫Steam类目发履约态变更
type AlibabaJymSteamFulfillmentUpdateAPIRequest struct {
	model.Params
	// 履约信息通知
	_deliveryNotifyDto *DeliveryNotifyDto
}

// NewAlibabaJymSteamFulfillmentUpdateRequest 初始化AlibabaJymSteamFulfillmentUpdateAPIRequest对象
func NewAlibabaJymSteamFulfillmentUpdateRequest() *AlibabaJymSteamFulfillmentUpdateAPIRequest {
	return &AlibabaJymSteamFulfillmentUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymSteamFulfillmentUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.steam.fulfillment.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymSteamFulfillmentUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymSteamFulfillmentUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryNotifyDto is DeliveryNotifyDto Setter
// 履约信息通知
func (r *AlibabaJymSteamFulfillmentUpdateAPIRequest) SetDeliveryNotifyDto(_deliveryNotifyDto *DeliveryNotifyDto) error {
	r._deliveryNotifyDto = _deliveryNotifyDto
	r.Set("delivery_notify_dto", _deliveryNotifyDto)
	return nil
}

// GetDeliveryNotifyDto DeliveryNotifyDto Getter
func (r AlibabaJymSteamFulfillmentUpdateAPIRequest) GetDeliveryNotifyDto() *DeliveryNotifyDto {
	return r._deliveryNotifyDto
}
