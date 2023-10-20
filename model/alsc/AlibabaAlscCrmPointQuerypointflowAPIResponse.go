package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointQuerypointflowAPIResponse 分页查询积分流水 API返回值
// alibaba.alsc.crm.point.querypointflow
//
// 分页查询积分流水
type AlibabaAlscCrmPointQuerypointflowAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPointQuerypointflowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPointQuerypointflowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmPointQuerypointflowAPIResponseModel).Reset()
}

// AlibabaAlscCrmPointQuerypointflowAPIResponseModel is 分页查询积分流水 成功返回结果
type AlibabaAlscCrmPointQuerypointflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_querypointflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPointQuerypointflowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmPointQuerypointflowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmPointQuerypointflowAPIResponse)
	},
}

// GetAlibabaAlscCrmPointQuerypointflowAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmPointQuerypointflowAPIResponse
func GetAlibabaAlscCrmPointQuerypointflowAPIResponse() *AlibabaAlscCrmPointQuerypointflowAPIResponse {
	return poolAlibabaAlscCrmPointQuerypointflowAPIResponse.Get().(*AlibabaAlscCrmPointQuerypointflowAPIResponse)
}

// ReleaseAlibabaAlscCrmPointQuerypointflowAPIResponse 将 AlibabaAlscCrmPointQuerypointflowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmPointQuerypointflowAPIResponse(v *AlibabaAlscCrmPointQuerypointflowAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmPointQuerypointflowAPIResponse.Put(v)
}
