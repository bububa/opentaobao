package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradeorderrefundlistqueryAPIRequest 查询退款单列表(卖家视角) API请求
// alibaba.lst.trade.order.refund.list.query
//
// 查询退款单列表(卖家视角)
type AlibabalsttradeorderrefundlistqueryAPIRequest struct {
	model.Params
	// 输入参数
	_param *TopLstSupplierOrderRefundQuery
}

// NewAlibabalsttradeorderrefundlistqueryRequest 初始化AlibabalsttradeorderrefundlistqueryAPIRequest对象
func NewAlibabalsttradeorderrefundlistqueryRequest() *AlibabalsttradeorderrefundlistqueryAPIRequest {
	return &AlibabalsttradeorderrefundlistqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradeorderrefundlistqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.refund.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradeorderrefundlistqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradeorderrefundlistqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 输入参数
func (r *AlibabalsttradeorderrefundlistqueryAPIRequest) SetParam(_param *TopLstSupplierOrderRefundQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabalsttradeorderrefundlistqueryAPIRequest) GetParam() *TopLstSupplierOrderRefundQuery {
	return r._param
}
