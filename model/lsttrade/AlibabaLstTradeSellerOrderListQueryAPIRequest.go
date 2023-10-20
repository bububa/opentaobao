package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerorderlistqueryAPIRequest 订单列表查看(卖家视角) API请求
// alibaba.lst.trade.seller.order.list.query
//
// 卖家视角订单查询，查询授权经销商订单列表
type AlibabalsttradesellerorderlistqueryAPIRequest struct {
	model.Params
	// 入参
	_param *LstTradeGetSellerOrderListParam
}

// NewAlibabalsttradesellerorderlistqueryRequest 初始化AlibabalsttradesellerorderlistqueryAPIRequest对象
func NewAlibabalsttradesellerorderlistqueryRequest() *AlibabalsttradesellerorderlistqueryAPIRequest {
	return &AlibabalsttradesellerorderlistqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradesellerorderlistqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.order.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradesellerorderlistqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradesellerorderlistqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabalsttradesellerorderlistqueryAPIRequest) SetParam(_param *LstTradeGetSellerOrderListParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabalsttradesellerorderlistqueryAPIRequest) GetParam() *LstTradeGetSellerOrderListParam {
	return r._param
}
