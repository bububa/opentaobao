package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest 销售活动绑定认购商品 API请求
// alibaba.alihouse.newhome.activity.subscription.bind
//
// 销售活动绑定认购商品
type AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest struct {
	model.Params
	// 入参属性值
	_subscriptionItemDto *SubscriptionItemDto
}

// NewAlibabaalihousenewhomeactivitysubscriptionbindRequest 初始化AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest对象
func NewAlibabaalihousenewhomeactivitysubscriptionbindRequest() *AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest {
	return &AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.subscription.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubscriptionItemDto is SubscriptionItemDto Setter
// 入参属性值
func (r *AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest) SetSubscriptionItemDto(_subscriptionItemDto *SubscriptionItemDto) error {
	r._subscriptionItemDto = _subscriptionItemDto
	r.Set("subscription_item_dto", _subscriptionItemDto)
	return nil
}

// GetSubscriptionItemDto SubscriptionItemDto Getter
func (r AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest) GetSubscriptionItemDto() *SubscriptionItemDto {
	return r._subscriptionItemDto
}
