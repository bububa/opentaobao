package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询物联网卡上订购的offer APIResponse
alibaba.aliqin.fc.iot.cardoffer

查询物联网卡上订购的offer
*/
type AlibabaAliqinFcIotCardofferAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcIotCardofferResponse `json:"alibaba_aliqin_fc_iot_cardoffer_response,omitempty"` 
    AlibabaAliqinFcIotCardofferResponse
}

/* model for simplify = false
type AlibabaAliqinFcIotCardofferResponse struct {

    // 结果对象
    
    Result  *struct {
        AlibabaAliqinFcIotCardofferResult  *AlibabaAliqinFcIotCardofferResult `json:"alibaba_aliqin_fc_iot_cardoffer_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcIotCardofferResponse struct {

    // 结果对象
    Result   *AlibabaAliqinFcIotCardofferResult `json:"result,omitempty"`

}
