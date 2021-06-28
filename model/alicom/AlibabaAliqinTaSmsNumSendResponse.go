package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
短信发送 APIResponse
alibaba.aliqin.ta.sms.num.send

短信发送
*/
type AlibabaAliqinTaSmsNumSendAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinTaSmsNumSendResponse `json:"alibaba_aliqin_ta_sms_num_send_response,omitempty"` 
    AlibabaAliqinTaSmsNumSendResponse
}

/* model for simplify = false
type AlibabaAliqinTaSmsNumSendResponse struct {

    // 返回结果
    
    Result  *struct {
        AlibabaAliqinTaSmsNumSendBizResult  *AlibabaAliqinTaSmsNumSendBizResult `json:"alibaba_aliqin_ta_sms_num_send_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinTaSmsNumSendResponse struct {

    // 返回结果
    Result   *AlibabaAliqinTaSmsNumSendBizResult `json:"result,omitempty"`

}
