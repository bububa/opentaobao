package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkWholesaleOrderCommitAPIRequest 盒马帮采购确认订单接口 API请求
// alibaba.wdk.wholesale.order.commit
//
// 盒马帮采购确认订单接口
type AlibabaWdkWholesaleOrderCommitAPIRequest struct {
	model.Params
	// 采购单信息
	_orderCommitReq *OrderCommitReq
}

// NewAlibabaWdkWholesaleOrderCommitRequest 初始化AlibabaWdkWholesaleOrderCommitAPIRequest对象
func NewAlibabaWdkWholesaleOrderCommitRequest() *AlibabaWdkWholesaleOrderCommitAPIRequest {
	return &AlibabaWdkWholesaleOrderCommitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkWholesaleOrderCommitAPIRequest) Reset() {
	r._orderCommitReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkWholesaleOrderCommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.wholesale.order.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkWholesaleOrderCommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkWholesaleOrderCommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCommitReq is OrderCommitReq Setter
// 采购单信息
func (r *AlibabaWdkWholesaleOrderCommitAPIRequest) SetOrderCommitReq(_orderCommitReq *OrderCommitReq) error {
	r._orderCommitReq = _orderCommitReq
	r.Set("order_commit_req", _orderCommitReq)
	return nil
}

// GetOrderCommitReq OrderCommitReq Getter
func (r AlibabaWdkWholesaleOrderCommitAPIRequest) GetOrderCommitReq() *OrderCommitReq {
	return r._orderCommitReq
}

var poolAlibabaWdkWholesaleOrderCommitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkWholesaleOrderCommitRequest()
	},
}

// GetAlibabaWdkWholesaleOrderCommitRequest 从 sync.Pool 获取 AlibabaWdkWholesaleOrderCommitAPIRequest
func GetAlibabaWdkWholesaleOrderCommitAPIRequest() *AlibabaWdkWholesaleOrderCommitAPIRequest {
	return poolAlibabaWdkWholesaleOrderCommitAPIRequest.Get().(*AlibabaWdkWholesaleOrderCommitAPIRequest)
}

// ReleaseAlibabaWdkWholesaleOrderCommitAPIRequest 将 AlibabaWdkWholesaleOrderCommitAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkWholesaleOrderCommitAPIRequest(v *AlibabaWdkWholesaleOrderCommitAPIRequest) {
	v.Reset()
	poolAlibabaWdkWholesaleOrderCommitAPIRequest.Put(v)
}
