package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinUserRegisterRemindAPIResponse isv到苗提醒 API返回值
// alibaba.health.vaccin.user.register.remind
//
// isv到苗提醒
type AlibabaHealthVaccinUserRegisterRemindAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinUserRegisterRemindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinUserRegisterRemindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinUserRegisterRemindAPIResponseModel).Reset()
}

// AlibabaHealthVaccinUserRegisterRemindAPIResponseModel is isv到苗提醒 成功返回结果
type AlibabaHealthVaccinUserRegisterRemindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_user_register_remind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// info
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 出参
	Model *IsvVcAvailableRemindResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinUserRegisterRemindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = nil
	m.IsSuccess = false
}

var poolAlibabaHealthVaccinUserRegisterRemindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinUserRegisterRemindAPIResponse)
	},
}

// GetAlibabaHealthVaccinUserRegisterRemindAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinUserRegisterRemindAPIResponse
func GetAlibabaHealthVaccinUserRegisterRemindAPIResponse() *AlibabaHealthVaccinUserRegisterRemindAPIResponse {
	return poolAlibabaHealthVaccinUserRegisterRemindAPIResponse.Get().(*AlibabaHealthVaccinUserRegisterRemindAPIResponse)
}

// ReleaseAlibabaHealthVaccinUserRegisterRemindAPIResponse 将 AlibabaHealthVaccinUserRegisterRemindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinUserRegisterRemindAPIResponse(v *AlibabaHealthVaccinUserRegisterRemindAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinUserRegisterRemindAPIResponse.Put(v)
}
