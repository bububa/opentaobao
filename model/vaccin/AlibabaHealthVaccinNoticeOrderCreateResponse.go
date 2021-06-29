package vaccin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝医疗健康疫苗预约创建 APIResponse
alibaba.health.vaccin.notice.order.create

支付宝医疗健康疫苗预约创建
*/
type AlibabaHealthVaccinNoticeOrderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaHealthVaccinNoticeOrderCreateResponse
}

type AlibabaHealthVaccinNoticeOrderCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_health_vaccin_notice_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回描述
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 返回码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 结果集
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
}
