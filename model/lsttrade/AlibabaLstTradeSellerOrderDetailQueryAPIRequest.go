package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOrderDetailQueryAPIRequest 订单详情查看(卖家视角) API请求
// alibaba.lst.trade.seller.order.detail.query
//
// 订单详情查看(卖家视角)
type AlibabaLstTradeSellerOrderDetailQueryAPIRequest struct {
	model.Params
	// 入参
	_param *LstTradeGetSellerOrderListParam
}

// NewAlibabaLstTradeSellerOrderDetailQueryRequest 初始化AlibabaLstTradeSellerOrderDetailQueryAPIRequest对象
func NewAlibabaLstTradeSellerOrderDetailQueryRequest() *AlibabaLstTradeSellerOrderDetailQueryAPIRequest {
	return &AlibabaLstTradeSellerOrderDetailQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOrderDetailQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.order.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOrderDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeSellerOrderDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaLstTradeSellerOrderDetailQueryAPIRequest) SetParam(_param *LstTradeGetSellerOrderListParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaLstTradeSellerOrderDetailQueryAPIRequest) GetParam() *LstTradeGetSellerOrderListParam {
	return r._param
}
