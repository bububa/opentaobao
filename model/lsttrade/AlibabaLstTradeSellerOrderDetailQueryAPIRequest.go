package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerorderdetailqueryAPIRequest 订单详情查看(卖家视角) API请求
// alibaba.lst.trade.seller.order.detail.query
//
// 订单详情查看(卖家视角)
type AlibabalsttradesellerorderdetailqueryAPIRequest struct {
	model.Params
	// 入参
	_param *LstTradeGetSellerOrderListParam
}

// NewAlibabalsttradesellerorderdetailqueryRequest 初始化AlibabalsttradesellerorderdetailqueryAPIRequest对象
func NewAlibabalsttradesellerorderdetailqueryRequest() *AlibabalsttradesellerorderdetailqueryAPIRequest {
	return &AlibabalsttradesellerorderdetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradesellerorderdetailqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.order.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradesellerorderdetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradesellerorderdetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabalsttradesellerorderdetailqueryAPIRequest) SetParam(_param *LstTradeGetSellerOrderListParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabalsttradesellerorderdetailqueryAPIRequest) GetParam() *LstTradeGetSellerOrderListParam {
	return r._param
}
