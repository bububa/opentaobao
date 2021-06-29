package vaccin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
福州疫苗取消预约 APIResponse
alibaba.health.vaccin.notice.order.cancel

福州疫苗用户取消预约接口
*/
type AlibabaHealthVaccinNoticeOrderCancelAPIResponse struct {
    model.CommonResponse
    AlibabaHealthVaccinNoticeOrderCancelResponse
}

type AlibabaHealthVaccinNoticeOrderCancelResponse struct {
    XMLName xml.Name `xml:"alibaba_health_vaccin_notice_order_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果集
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 返回码
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 返回描述
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
