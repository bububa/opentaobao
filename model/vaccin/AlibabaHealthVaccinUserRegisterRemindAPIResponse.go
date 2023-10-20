package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinuserregisterremindAPIResponse isv到苗提醒 API返回值
// alibaba.health.vaccin.user.register.remind
//
// isv到苗提醒
type AlibabahealthvaccinuserregisterremindAPIResponse struct {
	model.CommonResponse
	AlibabahealthvaccinuserregisterremindAPIResponseModel
}

// AlibabahealthvaccinuserregisterremindAPIResponseModel is isv到苗提醒 成功返回结果
type AlibabahealthvaccinuserregisterremindAPIResponseModel struct {
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
