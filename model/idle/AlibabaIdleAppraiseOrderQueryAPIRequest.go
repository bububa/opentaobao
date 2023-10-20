package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAppraiseOrderQueryAPIRequest 闲鱼验货宝订单详情查询 API请求
// alibaba.idle.appraise.order.query
//
// 鉴定商调用该接口获取订单状态
type AlibabaIdleAppraiseOrderQueryAPIRequest struct {
	model.Params
	// orderId
	_bizOrderId int64
}

// NewAlibabaIdleAppraiseOrderQueryRequest 初始化AlibabaIdleAppraiseOrderQueryAPIRequest对象
func NewAlibabaIdleAppraiseOrderQueryRequest() *AlibabaIdleAppraiseOrderQueryAPIRequest {
	return &AlibabaIdleAppraiseOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleAppraiseOrderQueryAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAppraiseOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAppraiseOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleAppraiseOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// orderId
func (r *AlibabaIdleAppraiseOrderQueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaIdleAppraiseOrderQueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolAlibabaIdleAppraiseOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleAppraiseOrderQueryRequest()
	},
}

// GetAlibabaIdleAppraiseOrderQueryRequest 从 sync.Pool 获取 AlibabaIdleAppraiseOrderQueryAPIRequest
func GetAlibabaIdleAppraiseOrderQueryAPIRequest() *AlibabaIdleAppraiseOrderQueryAPIRequest {
	return poolAlibabaIdleAppraiseOrderQueryAPIRequest.Get().(*AlibabaIdleAppraiseOrderQueryAPIRequest)
}

// ReleaseAlibabaIdleAppraiseOrderQueryAPIRequest 将 AlibabaIdleAppraiseOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleAppraiseOrderQueryAPIRequest(v *AlibabaIdleAppraiseOrderQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleAppraiseOrderQueryAPIRequest.Put(v)
}
