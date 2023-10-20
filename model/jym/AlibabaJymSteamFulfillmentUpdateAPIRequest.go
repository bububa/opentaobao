package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymsteamfulfillmentupdateAPIRequest 交易猫Steam类目发履约态变更 API请求
// alibaba.jym.steam.fulfillment.update
//
// 交易猫Steam类目发履约态变更
type AlibabajymsteamfulfillmentupdateAPIRequest struct {
	model.Params
	// 履约信息通知
	_deliveryNotifyDto *DeliveryNotifyDto
}

// NewAlibabajymsteamfulfillmentupdateRequest 初始化AlibabajymsteamfulfillmentupdateAPIRequest对象
func NewAlibabajymsteamfulfillmentupdateRequest() *AlibabajymsteamfulfillmentupdateAPIRequest {
	return &AlibabajymsteamfulfillmentupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymsteamfulfillmentupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.steam.fulfillment.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymsteamfulfillmentupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymsteamfulfillmentupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryNotifyDto is DeliveryNotifyDto Setter
// 履约信息通知
func (r *AlibabajymsteamfulfillmentupdateAPIRequest) SetDeliveryNotifyDto(_deliveryNotifyDto *DeliveryNotifyDto) error {
	r._deliveryNotifyDto = _deliveryNotifyDto
	r.Set("delivery_notify_dto", _deliveryNotifyDto)
	return nil
}

// GetDeliveryNotifyDto DeliveryNotifyDto Getter
func (r AlibabajymsteamfulfillmentupdateAPIRequest) GetDeliveryNotifyDto() *DeliveryNotifyDto {
	return r._deliveryNotifyDto
}
