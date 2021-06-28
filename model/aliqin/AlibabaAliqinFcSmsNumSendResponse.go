package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
短信发送 APIResponse
alibaba.aliqin.fc.sms.num.send

向指定手机号码发送模板短信，模板内可设置部分变量。使用前需要在阿里大于管理中心添加短信签名与短信模板。测试时请直接使用正式环境HTTP请求地址。
【重要】批量发送（一次传递多个号码eg:1381111111,1382222222）会产生相应的延迟，触达时间要求高的建议单条发送
*/
type AlibabaAliqinFcSmsNumSendAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcSmsNumSendResponse
}

type AlibabaAliqinFcSmsNumSendResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_sms_num_send_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *AlibabaAliqinFcSmsNumSendBizResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
