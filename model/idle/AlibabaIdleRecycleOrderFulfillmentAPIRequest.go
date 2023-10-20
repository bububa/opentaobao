package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerecycleorderfulfillmentAPIRequest 闲鱼回收订单履约V1 API请求
// alibaba.idle.recycle.order.fulfillment
//
// 外部回收商针对自有回收订单的履行
type AlibabaidlerecycleorderfulfillmentAPIRequest struct {
	model.Params
	// 订单同步入参
	_param0 *RecycleOrderSynDto
}

// NewAlibabaidlerecycleorderfulfillmentRequest 初始化AlibabaidlerecycleorderfulfillmentAPIRequest对象
func NewAlibabaidlerecycleorderfulfillmentRequest() *AlibabaidlerecycleorderfulfillmentAPIRequest {
	return &AlibabaidlerecycleorderfulfillmentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerecycleorderfulfillmentAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.fulfillment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerecycleorderfulfillmentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerecycleorderfulfillmentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 订单同步入参
func (r *AlibabaidlerecycleorderfulfillmentAPIRequest) SetParam0(_param0 *RecycleOrderSynDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaidlerecycleorderfulfillmentAPIRequest) GetParam0() *RecycleOrderSynDto {
	return r._param0
}
