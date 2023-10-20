package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest 快递送货上门能力同步/更新接口 API请求
// taobao.logistics.express.delivery.send.ability.async
//
// 快递送货上门能力同步/更新接口
type TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest struct {
	model.Params
	// 快递送货上门能力
	_deliverySendAbilityRequest *DeliverySendAbilityRequest
}

// NewTaobaoLogisticsExpressDeliverySendAbilityAsyncRequest 初始化TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest对象
func NewTaobaoLogisticsExpressDeliverySendAbilityAsyncRequest() *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest {
	return &TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.delivery.send.ability.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliverySendAbilityRequest is DeliverySendAbilityRequest Setter
// 快递送货上门能力
func (r *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) SetDeliverySendAbilityRequest(_deliverySendAbilityRequest *DeliverySendAbilityRequest) error {
	r._deliverySendAbilityRequest = _deliverySendAbilityRequest
	r.Set("delivery_send_ability_request", _deliverySendAbilityRequest)
	return nil
}

// GetDeliverySendAbilityRequest DeliverySendAbilityRequest Getter
func (r TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) GetDeliverySendAbilityRequest() *DeliverySendAbilityRequest {
	return r._deliverySendAbilityRequest
}
