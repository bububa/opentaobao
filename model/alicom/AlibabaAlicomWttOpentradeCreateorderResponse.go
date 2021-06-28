package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
充值送活动下单接口 APIResponse
alibaba.alicom.wtt.opentrade.createorder

提供给话费宝创建淘宝订单
*/
type AlibabaAlicomWttOpentradeCreateorderAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlicomWttOpentradeCreateorderResponse `json:"alibaba_alicom_wtt_opentrade_createorder_response,omitempty"` 
    AlibabaAlicomWttOpentradeCreateorderResponse
}

/* model for simplify = false
type AlibabaAlicomWttOpentradeCreateorderResponse struct {

    // result
    
    Result  *struct {
        TopResultDto  *TopResultDto `json:"top_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlicomWttOpentradeCreateorderResponse struct {

    // result
    Result   *TopResultDto `json:"result,omitempty"`

}
