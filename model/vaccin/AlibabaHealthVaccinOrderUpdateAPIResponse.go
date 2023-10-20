package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinOrderUpdateAPIResponse 回传取号状态 API返回值
// alibaba.health.vaccin.order.update
//
// 回传取号状态
type AlibabaHealthVaccinOrderUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinOrderUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinOrderUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinOrderUpdateAPIResponseModel).Reset()
}

// AlibabaHealthVaccinOrderUpdateAPIResponseModel is 回传取号状态 成功返回结果
type AlibabaHealthVaccinOrderUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_order_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 1
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 1
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 1
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinOrderUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
	m.IsSuccess = false
}

var poolAlibabaHealthVaccinOrderUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinOrderUpdateAPIResponse)
	},
}

// GetAlibabaHealthVaccinOrderUpdateAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinOrderUpdateAPIResponse
func GetAlibabaHealthVaccinOrderUpdateAPIResponse() *AlibabaHealthVaccinOrderUpdateAPIResponse {
	return poolAlibabaHealthVaccinOrderUpdateAPIResponse.Get().(*AlibabaHealthVaccinOrderUpdateAPIResponse)
}

// ReleaseAlibabaHealthVaccinOrderUpdateAPIResponse 将 AlibabaHealthVaccinOrderUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinOrderUpdateAPIResponse(v *AlibabaHealthVaccinOrderUpdateAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinOrderUpdateAPIResponse.Put(v)
}
