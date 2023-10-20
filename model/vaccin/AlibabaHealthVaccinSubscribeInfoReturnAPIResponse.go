package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinSubscribeInfoReturnAPIResponse 自有pov预约信息回传 API返回值
// alibaba.health.vaccin.subscribe.info.return
//
// 自有pov预约信息回传
type AlibabaHealthVaccinSubscribeInfoReturnAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinSubscribeInfoReturnAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinSubscribeInfoReturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinSubscribeInfoReturnAPIResponseModel).Reset()
}

// AlibabaHealthVaccinSubscribeInfoReturnAPIResponseModel is 自有pov预约信息回传 成功返回结果
type AlibabaHealthVaccinSubscribeInfoReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_subscribe_info_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// info
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *IsvPovSubscribeInfoResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinSubscribeInfoReturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = nil
	m.IsSuccess = false
}

var poolAlibabaHealthVaccinSubscribeInfoReturnAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinSubscribeInfoReturnAPIResponse)
	},
}

// GetAlibabaHealthVaccinSubscribeInfoReturnAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinSubscribeInfoReturnAPIResponse
func GetAlibabaHealthVaccinSubscribeInfoReturnAPIResponse() *AlibabaHealthVaccinSubscribeInfoReturnAPIResponse {
	return poolAlibabaHealthVaccinSubscribeInfoReturnAPIResponse.Get().(*AlibabaHealthVaccinSubscribeInfoReturnAPIResponse)
}

// ReleaseAlibabaHealthVaccinSubscribeInfoReturnAPIResponse 将 AlibabaHealthVaccinSubscribeInfoReturnAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinSubscribeInfoReturnAPIResponse(v *AlibabaHealthVaccinSubscribeInfoReturnAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinSubscribeInfoReturnAPIResponse.Put(v)
}
