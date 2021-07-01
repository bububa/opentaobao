package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleOrderFulfillmentAPIRequest
闲鱼回收订单履约V1 API请求
alibaba.idle.recycle.order.fulfillment

外部回收商针对自有回收订单的履行 */
type AlibabaIdleRecycleOrderFulfillmentAPIRequest struct {
	model.Params
	// 订单同步入参
	_param0 *RecycleOrderSynDto
}

// NewAlibabaIdleRecycleOrderFulfillmentRequest 初始化AlibabaIdleRecycleOrderFulfillmentAPIRequest对象
func NewAlibabaIdleRecycleOrderFulfillmentRequest() *AlibabaIdleRecycleOrderFulfillmentAPIRequest {
	return &AlibabaIdleRecycleOrderFulfillmentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderFulfillmentAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.fulfillment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderFulfillmentAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 订单同步入参
func (r *AlibabaIdleRecycleOrderFulfillmentAPIRequest) SetParam0(_param0 *RecycleOrderSynDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaIdleRecycleOrderFulfillmentAPIRequest) GetParam0() *RecycleOrderSynDto {
	return r._param0
}
