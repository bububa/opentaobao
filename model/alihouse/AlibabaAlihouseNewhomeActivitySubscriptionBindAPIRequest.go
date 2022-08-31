package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest 销售活动绑定认购商品 API请求
// alibaba.alihouse.newhome.activity.subscription.bind
//
// 销售活动绑定认购商品
type AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest struct {
	model.Params
	// 入参属性值
	_subscriptionItemDto *SubscriptionItemDto
}

// NewAlibabaAlihouseNewhomeActivitySubscriptionBindRequest 初始化AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest对象
func NewAlibabaAlihouseNewhomeActivitySubscriptionBindRequest() *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest {
	return &AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.subscription.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubscriptionItemDto is SubscriptionItemDto Setter
// 入参属性值
func (r *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) SetSubscriptionItemDto(_subscriptionItemDto *SubscriptionItemDto) error {
	r._subscriptionItemDto = _subscriptionItemDto
	r.Set("subscription_item_dto", _subscriptionItemDto)
	return nil
}

// GetSubscriptionItemDto SubscriptionItemDto Getter
func (r AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) GetSubscriptionItemDto() *SubscriptionItemDto {
	return r._subscriptionItemDto
}
