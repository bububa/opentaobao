package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressdeliverysendabilityasyncAPIRequest 快递送货上门能力同步/更新接口 API请求
// taobao.logistics.express.delivery.send.ability.async
//
// 快递送货上门能力同步/更新接口
type TaobaologisticsexpressdeliverysendabilityasyncAPIRequest struct {
	model.Params
	// 快递送货上门能力
	_deliverySendAbilityRequest *DeliverySendAbilityRequest
}

// NewTaobaologisticsexpressdeliverysendabilityasyncRequest 初始化TaobaologisticsexpressdeliverysendabilityasyncAPIRequest对象
func NewTaobaologisticsexpressdeliverysendabilityasyncRequest() *TaobaologisticsexpressdeliverysendabilityasyncAPIRequest {
	return &TaobaologisticsexpressdeliverysendabilityasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressdeliverysendabilityasyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.delivery.send.ability.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressdeliverysendabilityasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressdeliverysendabilityasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliverySendAbilityRequest is DeliverySendAbilityRequest Setter
// 快递送货上门能力
func (r *TaobaologisticsexpressdeliverysendabilityasyncAPIRequest) SetDeliverySendAbilityRequest(_deliverySendAbilityRequest *DeliverySendAbilityRequest) error {
	r._deliverySendAbilityRequest = _deliverySendAbilityRequest
	r.Set("delivery_send_ability_request", _deliverySendAbilityRequest)
	return nil
}

// GetDeliverySendAbilityRequest DeliverySendAbilityRequest Getter
func (r TaobaologisticsexpressdeliverysendabilityasyncAPIRequest) GetDeliverySendAbilityRequest() *DeliverySendAbilityRequest {
	return r._deliverySendAbilityRequest
}
