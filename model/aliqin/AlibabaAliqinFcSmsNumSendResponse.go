package aliqin

import (
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
    // Response *AlibabaAliqinFcSmsNumSendResponse `json:"alibaba_aliqin_fc_sms_num_send_response,omitempty"` 
    AlibabaAliqinFcSmsNumSendResponse
}

/* model for simplify = false
type AlibabaAliqinFcSmsNumSendResponse struct {

    // 返回值
    
    Result  *struct {
        AlibabaAliqinFcSmsNumSendBizResult  *AlibabaAliqinFcSmsNumSendBizResult `json:"alibaba_aliqin_fc_sms_num_send_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcSmsNumSendResponse struct {

    // 返回值
    Result   *AlibabaAliqinFcSmsNumSendBizResult `json:"result,omitempty"`

}
