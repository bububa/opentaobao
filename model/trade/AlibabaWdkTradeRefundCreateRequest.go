package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道逆向订单创建 API请求
alibaba.wdk.trade.refund.create

该接口是创建退货订单的服务。当外部渠道发起退款后，调用此接口可以完成五道口底层交易、履约、配送等一系列流程进行退货。
*/
type AlibabaWdkTradeRefundCreateRequest struct {
    model.Params
    // 退货请求
    _refundGoodsCreateRequest   *RefundGoodsCreateRequest
}

// 初始化AlibabaWdkTradeRefundCreateRequest对象
func NewAlibabaWdkTradeRefundCreateRequest() *AlibabaWdkTradeRefundCreateRequest{
    return &AlibabaWdkTradeRefundCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeRefundCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundGoodsCreateRequest Setter
// 退货请求
func (r *AlibabaWdkTradeRefundCreateRequest) SetRefundGoodsCreateRequest(_refundGoodsCreateRequest *RefundGoodsCreateRequest) error {
    r._refundGoodsCreateRequest = _refundGoodsCreateRequest
    r.Set("refund_goods_create_request", _refundGoodsCreateRequest)
    return nil
}

// RefundGoodsCreateRequest Getter
func (r AlibabaWdkTradeRefundCreateRequest) GetRefundGoodsCreateRequest() *RefundGoodsCreateRequest {
    return r._refundGoodsCreateRequest
}
