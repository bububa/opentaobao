package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinPovUpdateAPIResponse 新增/变更接种点信息 API返回值
// alibaba.health.vaccin.pov.update
//
// ISV 将疫苗的接种点信息同步到免疫规划中心，提醒用户接种时可提供接种点详情。
type AlibabaHealthVaccinPovUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinPovUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinPovUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinPovUpdateAPIResponseModel).Reset()
}

// AlibabaHealthVaccinPovUpdateAPIResponseModel is 新增/变更接种点信息 成功返回结果
type AlibabaHealthVaccinPovUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_pov_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回数据详情
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功执行
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinPovUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
	m.IsSuccess = false
}

var poolAlibabaHealthVaccinPovUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinPovUpdateAPIResponse)
	},
}

// GetAlibabaHealthVaccinPovUpdateAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinPovUpdateAPIResponse
func GetAlibabaHealthVaccinPovUpdateAPIResponse() *AlibabaHealthVaccinPovUpdateAPIResponse {
	return poolAlibabaHealthVaccinPovUpdateAPIResponse.Get().(*AlibabaHealthVaccinPovUpdateAPIResponse)
}

// ReleaseAlibabaHealthVaccinPovUpdateAPIResponse 将 AlibabaHealthVaccinPovUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinPovUpdateAPIResponse(v *AlibabaHealthVaccinPovUpdateAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinPovUpdateAPIResponse.Put(v)
}
