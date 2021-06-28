package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外部渠道通知淘鲜达退款成功接口 APIResponse
alibaba.wdk.trade.refund.inform

该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
*/
type AlibabaWdkTradeRefundInformAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkTradeRefundInformResponse `json:"alibaba_wdk_trade_refund_inform_response,omitempty"` 
    AlibabaWdkTradeRefundInformResponse
}

/* model for simplify = false
type AlibabaWdkTradeRefundInformResponse struct {

    // 返回结果
    
    Result  *struct {
        InformRefundSuccessResult  *InformRefundSuccessResult `json:"inform_refund_success_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkTradeRefundInformResponse struct {

    // 返回结果
    Result   *InformRefundSuccessResult `json:"result,omitempty"`

}
