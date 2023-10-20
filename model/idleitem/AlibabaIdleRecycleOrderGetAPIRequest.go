package idleitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderGetAPIRequest 闲鱼回收订单查询V2 API请求
// alibaba.idle.recycle.order.get
//
// 闲鱼回收业务中,外部回收商作为交易上买家,闲鱼用户下单后,需要回收商主动拉取交易订单
type AlibabaIdleRecycleOrderGetAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewAlibabaIdleRecycleOrderGetRequest 初始化AlibabaIdleRecycleOrderGetAPIRequest对象
func NewAlibabaIdleRecycleOrderGetRequest() *AlibabaIdleRecycleOrderGetAPIRequest {
	return &AlibabaIdleRecycleOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRecycleOrderGetAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRecycleOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *AlibabaIdleRecycleOrderGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaIdleRecycleOrderGetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolAlibabaIdleRecycleOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRecycleOrderGetRequest()
	},
}

// GetAlibabaIdleRecycleOrderGetRequest 从 sync.Pool 获取 AlibabaIdleRecycleOrderGetAPIRequest
func GetAlibabaIdleRecycleOrderGetAPIRequest() *AlibabaIdleRecycleOrderGetAPIRequest {
	return poolAlibabaIdleRecycleOrderGetAPIRequest.Get().(*AlibabaIdleRecycleOrderGetAPIRequest)
}

// ReleaseAlibabaIdleRecycleOrderGetAPIRequest 将 AlibabaIdleRecycleOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRecycleOrderGetAPIRequest(v *AlibabaIdleRecycleOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaIdleRecycleOrderGetAPIRequest.Put(v)
}
