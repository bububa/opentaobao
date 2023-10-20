package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDruguseQueryAPIResponse 合理用药规则查询 API返回值
// alibaba.alihealth.druguse.query
//
// 查询用户购买的药品命中的风险规则
type AlibabaAlihealthDruguseQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDruguseQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDruguseQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDruguseQueryAPIResponseModel).Reset()
}

// AlibabaAlihealthDruguseQueryAPIResponseModel is 合理用药规则查询 成功返回结果
type AlibabaAlihealthDruguseQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_druguse_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDruguseQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDruguseQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDruguseQueryAPIResponse)
	},
}

// GetAlibabaAlihealthDruguseQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDruguseQueryAPIResponse
func GetAlibabaAlihealthDruguseQueryAPIResponse() *AlibabaAlihealthDruguseQueryAPIResponse {
	return poolAlibabaAlihealthDruguseQueryAPIResponse.Get().(*AlibabaAlihealthDruguseQueryAPIResponse)
}

// ReleaseAlibabaAlihealthDruguseQueryAPIResponse 将 AlibabaAlihealthDruguseQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDruguseQueryAPIResponse(v *AlibabaAlihealthDruguseQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDruguseQueryAPIResponse.Put(v)
}
