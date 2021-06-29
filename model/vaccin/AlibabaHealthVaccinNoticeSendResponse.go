package vaccin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发送消息提醒 APIResponse
alibaba.health.vaccin.notice.send

ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。
*/
type AlibabaHealthVaccinNoticeSendAPIResponse struct {
    model.CommonResponse
    AlibabaHealthVaccinNoticeSendResponse
}

type AlibabaHealthVaccinNoticeSendResponse struct {
    XMLName xml.Name `xml:"alibaba_health_vaccin_notice_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功执行
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 有数据返回时的数据详情
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 找不到疫苗信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 200
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
