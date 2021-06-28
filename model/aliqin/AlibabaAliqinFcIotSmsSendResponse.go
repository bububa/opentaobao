package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物联网短信发送 APIResponse
alibaba.aliqin.fc.iot.sms.send

发送物联网短信，只允许使用物联网短信模板
*/
type AlibabaAliqinFcIotSmsSendAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcIotSmsSendResponse `json:"alibaba_aliqin_fc_iot_sms_send_response,omitempty"` 
    AlibabaAliqinFcIotSmsSendResponse
}

/* model for simplify = false
type AlibabaAliqinFcIotSmsSendResponse struct {

    // 返回值
    
    Result  *struct {
        AlibabaAliqinFcIotSmsSendBizResult  *AlibabaAliqinFcIotSmsSendBizResult `json:"alibaba_aliqin_fc_iot_sms_send_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcIotSmsSendResponse struct {

    // 返回值
    Result   *AlibabaAliqinFcIotSmsSendBizResult `json:"result,omitempty"`

}
