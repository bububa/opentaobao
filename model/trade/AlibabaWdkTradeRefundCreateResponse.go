package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道逆向订单创建 APIResponse
alibaba.wdk.trade.refund.create

该接口是创建退货订单的服务。当外部渠道发起退款后，调用此接口可以完成五道口底层交易、履约、配送等一系列流程进行退货。
*/
type AlibabaWdkTradeRefundCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkTradeRefundCreateResponse `json:"alibaba_wdk_trade_refund_create_response,omitempty"`
}

type AlibabaWdkTradeRefundCreateResponse struct {

    // 返回结果
    Result   *RefundGoodsCreateResult `json:"result,omitempty"`

}
