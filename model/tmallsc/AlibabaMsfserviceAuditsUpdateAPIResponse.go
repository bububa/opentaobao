package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMsfserviceAuditsUpdateAPIResponse 操作改约审批单 API返回值
// alibaba.msfservice.audits.update
//
// 操作改约审批单
type AlibabaMsfserviceAuditsUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaMsfserviceAuditsUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMsfserviceAuditsUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMsfserviceAuditsUpdateAPIResponseModel).Reset()
}

// AlibabaMsfserviceAuditsUpdateAPIResponseModel is 操作改约审批单 成功返回结果
type AlibabaMsfserviceAuditsUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_msfservice_audits_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlibabaMsfserviceAuditsUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMsfserviceAuditsUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMsfserviceAuditsUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMsfserviceAuditsUpdateAPIResponse)
	},
}

// GetAlibabaMsfserviceAuditsUpdateAPIResponse 从 sync.Pool 获取 AlibabaMsfserviceAuditsUpdateAPIResponse
func GetAlibabaMsfserviceAuditsUpdateAPIResponse() *AlibabaMsfserviceAuditsUpdateAPIResponse {
	return poolAlibabaMsfserviceAuditsUpdateAPIResponse.Get().(*AlibabaMsfserviceAuditsUpdateAPIResponse)
}

// ReleaseAlibabaMsfserviceAuditsUpdateAPIResponse 将 AlibabaMsfserviceAuditsUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMsfserviceAuditsUpdateAPIResponse(v *AlibabaMsfserviceAuditsUpdateAPIResponse) {
	v.Reset()
	poolAlibabaMsfserviceAuditsUpdateAPIResponse.Put(v)
}
