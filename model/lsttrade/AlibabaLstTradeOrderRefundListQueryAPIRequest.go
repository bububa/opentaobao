package lsttrade

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderRefundListQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.refund.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderRefundListQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
