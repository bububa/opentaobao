package lsttrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeOrderRefundListQueryAPIRequest 查询退款单列表(卖家视角) API请求
// alibaba.lst.trade.order.refund.list.query
//
// 查询退款单列表(卖家视角)
type AlibabaLstTradeOrderRefundListQueryAPIRequest struct {
	model.Params
	// 输入参数
	_param *TopLstSupplierOrderRefundQuery
}

// NewAlibabaLstTradeOrderRefundListQueryRequest 初始化AlibabaLstTradeOrderRefundListQueryAPIRequest对象
func NewAlibabaLstTradeOrderRefundListQueryRequest() *AlibabaLstTradeOrderRefundListQueryAPIRequest {
	return &AlibabaLstTradeOrderRefundListQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeOrderRefundListQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderRefundListQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.refund.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderRefundListQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeOrderRefundListQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 输入参数
func (r *AlibabaLstTradeOrderRefundListQueryAPIRequest) SetParam(_param *TopLstSupplierOrderRefundQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaLstTradeOrderRefundListQueryAPIRequest) GetParam() *TopLstSupplierOrderRefundQuery {
	return r._param
}

var poolAlibabaLstTradeOrderRefundListQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeOrderRefundListQueryRequest()
	},
}

// GetAlibabaLstTradeOrderRefundListQueryRequest 从 sync.Pool 获取 AlibabaLstTradeOrderRefundListQueryAPIRequest
func GetAlibabaLstTradeOrderRefundListQueryAPIRequest() *AlibabaLstTradeOrderRefundListQueryAPIRequest {
	return poolAlibabaLstTradeOrderRefundListQueryAPIRequest.Get().(*AlibabaLstTradeOrderRefundListQueryAPIRequest)
}

// ReleaseAlibabaLstTradeOrderRefundListQueryAPIRequest 将 AlibabaLstTradeOrderRefundListQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeOrderRefundListQueryAPIRequest(v *AlibabaLstTradeOrderRefundListQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeOrderRefundListQueryAPIRequest.Put(v)
}
