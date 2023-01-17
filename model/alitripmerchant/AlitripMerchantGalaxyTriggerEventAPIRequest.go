package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyTriggerEventAPIRequest 抽奖活动自定义事件触发 API请求
// alitrip.merchant.galaxy.trigger.event
//
// 自定义事件触发
type AlitripMerchantGalaxyTriggerEventAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 触发事件入参
	_eventParam *EventParam
}

// NewAlitripMerchantGalaxyTriggerEventRequest 初始化AlitripMerchantGalaxyTriggerEventAPIRequest对象
func NewAlitripMerchantGalaxyTriggerEventRequest() *AlitripMerchantGalaxyTriggerEventAPIRequest {
	return &AlitripMerchantGalaxyTriggerEventAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyTriggerEventAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.trigger.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyTriggerEventAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyTriggerEventAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripMerchantGalaxyTriggerEventAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyTriggerEventAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetEventParam is EventParam Setter
// 触发事件入参
func (r *AlitripMerchantGalaxyTriggerEventAPIRequest) SetEventParam(_eventParam *EventParam) error {
	r._eventParam = _eventParam
	r.Set("event_param", _eventParam)
	return nil
}

// GetEventParam EventParam Getter
func (r AlitripMerchantGalaxyTriggerEventAPIRequest) GetEventParam() *EventParam {
	return r._eventParam
}
