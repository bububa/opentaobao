package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktraderefundqueryAPIRequest 外部渠道查询退货订单详情接口 API请求
// alibaba.wdk.trade.refund.query
//
// 该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。
type AlibabawdktraderefundqueryAPIRequest struct {
	model.Params
	// 查询请求
	_refundGoodsQueryRequest *RefundGoodsQueryRequest
}

// NewAlibabawdktraderefundqueryRequest 初始化AlibabawdktraderefundqueryAPIRequest对象
func NewAlibabawdktraderefundqueryRequest() *AlibabawdktraderefundqueryAPIRequest {
	return &AlibabawdktraderefundqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktraderefundqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktraderefundqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktraderefundqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundGoodsQueryRequest is RefundGoodsQueryRequest Setter
// 查询请求
func (r *AlibabawdktraderefundqueryAPIRequest) SetRefundGoodsQueryRequest(_refundGoodsQueryRequest *RefundGoodsQueryRequest) error {
	r._refundGoodsQueryRequest = _refundGoodsQueryRequest
	r.Set("refund_goods_query_request", _refundGoodsQueryRequest)
	return nil
}

// GetRefundGoodsQueryRequest RefundGoodsQueryRequest Getter
func (r AlibabawdktraderefundqueryAPIRequest) GetRefundGoodsQueryRequest() *RefundGoodsQueryRequest {
	return r._refundGoodsQueryRequest
}
