package vaccin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约时间段提醒 APIResponse
alibaba.health.vaccin.notice.timebucket.remind

疫苗预约时间段提醒
*/
type AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse struct {
    model.CommonResponse
    AlibabaHealthVaccinNoticeTimebucketRemindResponse
}

type AlibabaHealthVaccinNoticeTimebucketRemindResponse struct {
    XMLName xml.Name `xml:"alibaba_health_vaccin_notice_timebucket_remind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 执行是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误code
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 错误描述
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
