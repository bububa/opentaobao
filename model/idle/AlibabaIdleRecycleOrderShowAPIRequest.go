package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderShowAPIRequest 闲鱼回收订单查询V1.1 API请求
// alibaba.idle.recycle.order.show
//
// 查询回收订单
type AlibabaIdleRecycleOrderShowAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewAlibabaIdleRecycleOrderShowRequest 初始化AlibabaIdleRecycleOrderShowAPIRequest对象
func NewAlibabaIdleRecycleOrderShowRequest() *AlibabaIdleRecycleOrderShowAPIRequest {
	return &AlibabaIdleRecycleOrderShowAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRecycleOrderShowAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderShowAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.show"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderShowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRecycleOrderShowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *AlibabaIdleRecycleOrderShowAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaIdleRecycleOrderShowAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolAlibabaIdleRecycleOrderShowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRecycleOrderShowRequest()
	},
}

// GetAlibabaIdleRecycleOrderShowRequest 从 sync.Pool 获取 AlibabaIdleRecycleOrderShowAPIRequest
func GetAlibabaIdleRecycleOrderShowAPIRequest() *AlibabaIdleRecycleOrderShowAPIRequest {
	return poolAlibabaIdleRecycleOrderShowAPIRequest.Get().(*AlibabaIdleRecycleOrderShowAPIRequest)
}

// ReleaseAlibabaIdleRecycleOrderShowAPIRequest 将 AlibabaIdleRecycleOrderShowAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRecycleOrderShowAPIRequest(v *AlibabaIdleRecycleOrderShowAPIRequest) {
	v.Reset()
	poolAlibabaIdleRecycleOrderShowAPIRequest.Put(v)
}
