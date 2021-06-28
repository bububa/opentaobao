package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单创建接口 APIResponse
alibaba.omni.saas.order.create

服务商利用现有的saas系统和阿里完成交易系统的对接
*/
type AlibabaOmniSaasOrderCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaOmniSaasOrderCreateResponse `json:"alibaba_omni_saas_order_create_response,omitempty"` 
    AlibabaOmniSaasOrderCreateResponse
}

/* model for simplify = false
type AlibabaOmniSaasOrderCreateResponse struct {

    // tradeNo
    
    TradeNo   int64 `json:"trade_no,omitempty"`
    

    // totalAmount
    
    TotalAmount   int64 `json:"total_amount,omitempty"`
    

    // actualPayFee
    
    ActualPayFee   int64 `json:"actual_pay_fee,omitempty"`
    

}
*/

type AlibabaOmniSaasOrderCreateResponse struct {

    // tradeNo
    TradeNo   int64 `json:"trade_no,omitempty"`

    // totalAmount
    TotalAmount   int64 `json:"total_amount,omitempty"`

    // actualPayFee
    ActualPayFee   int64 `json:"actual_pay_fee,omitempty"`

}
