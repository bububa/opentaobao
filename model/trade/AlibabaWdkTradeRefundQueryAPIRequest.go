package trade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTradeRefundQueryAPIRequest) Reset() {
	r._refundGoodsQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeRefundQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeRefundQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeRefundQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkTradeRefundQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTradeRefundQueryRequest()
	},
}

// GetAlibabaWdkTradeRefundQueryRequest 从 sync.Pool 获取 AlibabaWdkTradeRefundQueryAPIRequest
func GetAlibabaWdkTradeRefundQueryAPIRequest() *AlibabaWdkTradeRefundQueryAPIRequest {
	return poolAlibabaWdkTradeRefundQueryAPIRequest.Get().(*AlibabaWdkTradeRefundQueryAPIRequest)
}

// ReleaseAlibabaWdkTradeRefundQueryAPIRequest 将 AlibabaWdkTradeRefundQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTradeRefundQueryAPIRequest(v *AlibabaWdkTradeRefundQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkTradeRefundQueryAPIRequest.Put(v)
}
