package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeRefundQueryAPIRequest 外部渠道查询退货订单详情接口 API请求
// alibaba.wdk.trade.refund.query
//
// 该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。
type AlibabaWdkTradeRefundQueryAPIRequest struct {
	model.Params
	// 查询请求
	_refundGoodsQueryRequest *RefundGoodsQueryRequest
}

// NewAlibabaWdkTradeRefundQueryRequest 初始化AlibabaWdkTradeRefundQueryAPIRequest对象
func NewAlibabaWdkTradeRefundQueryRequest() *AlibabaWdkTradeRefundQueryAPIRequest {
	return &AlibabaWdkTradeRefundQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeRefundQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeRefundQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefundGoodsQueryRequest is RefundGoodsQueryRequest Setter
// 查询请求
func (r *AlibabaWdkTradeRefundQueryAPIRequest) SetRefundGoodsQueryRequest(_refundGoodsQueryRequest *RefundGoodsQueryRequest) error {
	r._refundGoodsQueryRequest = _refundGoodsQueryRequest
	r.Set("refund_goods_query_request", _refundGoodsQueryRequest)
	return nil
}

// GetRefundGoodsQueryRequest RefundGoodsQueryRequest Getter
func (r AlibabaWdkTradeRefundQueryAPIRequest) GetRefundGoodsQueryRequest() *RefundGoodsQueryRequest {
	return r._refundGoodsQueryRequest
}
