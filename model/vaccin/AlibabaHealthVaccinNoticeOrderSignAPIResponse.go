package vaccin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
福州疫苗签到成功通知 API返回值 
alibaba.health.vaccin.notice.order.sign

福州疫苗用户签到成功记录
*/
type AlibabaHealthVaccinNoticeOrderSignAPIResponse struct {
    model.CommonResponse
    AlibabaHealthVaccinNoticeOrderSignAPIResponseModel
}

// 福州疫苗签到成功通知 成功返回结果
type AlibabaHealthVaccinNoticeOrderSignAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_health_vaccin_notice_order_sign_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果集
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // 返回码
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 返回描述
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
