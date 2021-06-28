package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按终端号订购增值业务 APIResponse
alibaba.aliqin.fc.iot.rechargeCard

按终端号订购增值业务
*/
type AlibabaAliqinFcIotRechargeCardAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcIotRechargeCardResponse `json:"alibaba_aliqin_fc_iot_rechargeCard_response,omitempty"` 
    AlibabaAliqinFcIotRechargeCardResponse
}

/* model for simplify = false
type AlibabaAliqinFcIotRechargeCardResponse struct {

    // result
    
    Result  *struct {
        AlibabaAliqinFcIotRechargeCardResult  *AlibabaAliqinFcIotRechargeCardResult `json:"alibaba_aliqin_fc_iot_recharge_card_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcIotRechargeCardResponse struct {

    // result
    Result   *AlibabaAliqinFcIotRechargeCardResult `json:"result,omitempty"`

}
