package seaking

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingServicepackAPIResponse 获取海王用户权限包 API返回值
// alibaba.seaking.servicepack
//
// 获取海王用户权限包
type AlibabaSeakingServicepackAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingServicepackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSeakingServicepackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingServicepackAPIResponseModel).Reset()
}

// AlibabaSeakingServicepackAPIResponseModel is 获取海王用户权限包 成功返回结果
type AlibabaSeakingServicepackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_servicepack_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 权限包列表
	ServicePackList []AlibabaSeakingServicepackResult `json:"service_pack_list,omitempty" xml:"service_pack_list>alibaba_seaking_servicepack_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingServicepackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServicePackList = m.ServicePackList[:0]
}

var poolAlibabaSeakingServicepackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingServicepackAPIResponse)
	},
}

// GetAlibabaSeakingServicepackAPIResponse 从 sync.Pool 获取 AlibabaSeakingServicepackAPIResponse
func GetAlibabaSeakingServicepackAPIResponse() *AlibabaSeakingServicepackAPIResponse {
	return poolAlibabaSeakingServicepackAPIResponse.Get().(*AlibabaSeakingServicepackAPIResponse)
}

// ReleaseAlibabaSeakingServicepackAPIResponse 将 AlibabaSeakingServicepackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingServicepackAPIResponse(v *AlibabaSeakingServicepackAPIResponse) {
	v.Reset()
	poolAlibabaSeakingServicepackAPIResponse.Put(v)
}
