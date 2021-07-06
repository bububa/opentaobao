package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeUserBindAPIResponse 支付宝疫苗绑定接种人 API返回值
// alibaba.health.vaccin.notice.user.bind
//
// 支付宝疫苗绑定接种人
type AlibabaHealthVaccinNoticeUserBindAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeUserBindAPIResponseModel
}

// AlibabaHealthVaccinNoticeUserBindAPIResponseModel is 支付宝疫苗绑定接种人 成功返回结果
type AlibabaHealthVaccinNoticeUserBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_user_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 执行是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
