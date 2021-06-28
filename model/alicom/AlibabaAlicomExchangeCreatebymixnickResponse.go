package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口tae APIResponse
alibaba.alicom.exchange.createbymixnick

代理商调用该接口来进行积分兑换，tae
*/
type AlibabaAlicomExchangeCreatebymixnickAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlicomExchangeCreatebymixnickResponse `json:"alibaba_alicom_exchange_createbymixnick_response,omitempty"` 
    AlibabaAlicomExchangeCreatebymixnickResponse
}

/* model for simplify = false
type AlibabaAlicomExchangeCreatebymixnickResponse struct {

    // result
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlicomExchangeCreatebymixnickResponse struct {

    // result
    Result   *CommonResult `json:"result,omitempty"`

}
