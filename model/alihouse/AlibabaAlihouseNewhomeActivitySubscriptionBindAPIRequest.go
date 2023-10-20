package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) Reset() {
	r._subscriptionItemDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.subscription.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeActivitySubscriptionBindRequest()
	},
}

// GetAlibabaAlihouseNewhomeActivitySubscriptionBindRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest
func GetAlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest() *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest {
	return poolAlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest.Get().(*AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest 将 AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest(v *AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest.Put(v)
}
