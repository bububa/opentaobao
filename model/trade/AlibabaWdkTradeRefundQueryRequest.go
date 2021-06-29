package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道查询退货订单详情接口 API请求
alibaba.wdk.trade.refund.query

该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。
*/
type AlibabaWdkTradeRefundQueryRequest struct {
    model.Params
    // 查询请求
    _refundGoodsQueryRequest   *RefundGoodsQueryRequest
}

// 初始化AlibabaWdkTradeRefundQueryRequest对象
func NewAlibabaWdkTradeRefundQueryRequest() *AlibabaWdkTradeRefundQueryRequest{
    return &AlibabaWdkTradeRefundQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeRefundQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.refund.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeRefundQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundGoodsQueryRequest Setter
// 查询请求
func (r *AlibabaWdkTradeRefundQueryRequest) SetRefundGoodsQueryRequest(_refundGoodsQueryRequest *RefundGoodsQueryRequest) error {
    r._refundGoodsQueryRequest = _refundGoodsQueryRequest
    r.Set("refund_goods_query_request", _refundGoodsQueryRequest)
    return nil
}

// RefundGoodsQueryRequest Getter
func (r AlibabaWdkTradeRefundQueryRequest) GetRefundGoodsQueryRequest() *RefundGoodsQueryRequest {
    return r._refundGoodsQueryRequest
}
