package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeUserCreateAPIResponse 支付宝医疗健康疫苗用户创建 API返回值
// alibaba.health.vaccin.notice.user.create
//
// 支付宝医疗健康疫苗用户创建
type AlibabaHealthVaccinNoticeUserCreateAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeUserCreateAPIResponseModel
}

// AlibabaHealthVaccinNoticeUserCreateAPIResponseModel is 支付宝医疗健康疫苗用户创建 成功返回结果
type AlibabaHealthVaccinNoticeUserCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_user_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 结果集
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
