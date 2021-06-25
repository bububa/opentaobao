package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道逆向订单创建 APIRequest
alibaba.wdk.trade.refund.create

该接口是创建退货订单的服务。当外部渠道发起退款后，调用此接口可以完成五道口底层交易、履约、配送等一系列流程进行退货。
*/
type AlibabaWdkTradeRefundCreateRequest struct {
    model.Params

    // 退货请求
    refundGoodsCreateRequest   *RefundGoodsCreateRequest 

}

func NewAlibabaWdkTradeRefundCreateRequest() *AlibabaWdkTradeRefundCreateRequest{
    return &AlibabaWdkTradeRefundCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTradeRefundCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.refund.create"
}

func (r AlibabaWdkTradeRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkTradeRefundCreateRequest) SetRefundGoodsCreateRequest(refundGoodsCreateRequest *RefundGoodsCreateRequest) error {
    r.refundGoodsCreateRequest = refundGoodsCreateRequest
    r.Set("refund_goods_create_request", refundGoodsCreateRequest)
    return nil
}

func (r AlibabaWdkTradeRefundCreateRequest) GetRefundGoodsCreateRequest() *RefundGoodsCreateRequest {
    return r.refundGoodsCreateRequest
}

