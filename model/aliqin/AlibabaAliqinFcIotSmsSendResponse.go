package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物联网短信发送 APIResponse
alibaba.aliqin.fc.iot.sms.send

发送物联网短信，只允许使用物联网短信模板
*/
type AlibabaAliqinFcIotSmsSendAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotSmsSendResponse
}

type AlibabaAliqinFcIotSmsSendResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_sms_send_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *AlibabaAliqinFcIotSmsSendBizResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
