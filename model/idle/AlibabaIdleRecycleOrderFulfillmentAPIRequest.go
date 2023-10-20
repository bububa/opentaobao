package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderFulfillmentAPIRequest 闲鱼回收订单履约V1 API请求
// alibaba.idle.recycle.order.fulfillment
//
// 外部回收商针对自有回收订单的履行
type AlibabaIdleRecycleOrderFulfillmentAPIRequest struct {
	model.Params
	// 订单同步入参
	_param0 *RecycleOrderSynDto
}

// NewAlibabaIdleRecycleOrderFulfillmentRequest 初始化AlibabaIdleRecycleOrderFulfillmentAPIRequest对象
func NewAlibabaIdleRecycleOrderFulfillmentRequest() *AlibabaIdleRecycleOrderFulfillmentAPIRequest {
	return &AlibabaIdleRecycleOrderFulfillmentAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRecycleOrderFulfillmentAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderFulfillmentAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.fulfillment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderFulfillmentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRecycleOrderFulfillmentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 订单同步入参
func (r *AlibabaIdleRecycleOrderFulfillmentAPIRequest) SetParam0(_param0 *RecycleOrderSynDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaIdleRecycleOrderFulfillmentAPIRequest) GetParam0() *RecycleOrderSynDto {
	return r._param0
}

var poolAlibabaIdleRecycleOrderFulfillmentAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRecycleOrderFulfillmentRequest()
	},
}

// GetAlibabaIdleRecycleOrderFulfillmentRequest 从 sync.Pool 获取 AlibabaIdleRecycleOrderFulfillmentAPIRequest
func GetAlibabaIdleRecycleOrderFulfillmentAPIRequest() *AlibabaIdleRecycleOrderFulfillmentAPIRequest {
	return poolAlibabaIdleRecycleOrderFulfillmentAPIRequest.Get().(*AlibabaIdleRecycleOrderFulfillmentAPIRequest)
}

// ReleaseAlibabaIdleRecycleOrderFulfillmentAPIRequest 将 AlibabaIdleRecycleOrderFulfillmentAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRecycleOrderFulfillmentAPIRequest(v *AlibabaIdleRecycleOrderFulfillmentAPIRequest) {
	v.Reset()
	poolAlibabaIdleRecycleOrderFulfillmentAPIRequest.Put(v)
}
