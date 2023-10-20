package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTradeitemAPIResponse 新二租品同步 API返回值
// alibaba.alihouse.newhome.project.tradeitem
//
// 新二品同步
type AlibabaAlihouseNewhomeProjectTradeitemAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectTradeitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectTradeitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectTradeitemAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectTradeitemAPIResponseModel is 新二租品同步 成功返回结果
type AlibabaAlihouseNewhomeProjectTradeitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_tradeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectTradeitemResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectTradeitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectTradeitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectTradeitemAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectTradeitemAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectTradeitemAPIResponse
func GetAlibabaAlihouseNewhomeProjectTradeitemAPIResponse() *AlibabaAlihouseNewhomeProjectTradeitemAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectTradeitemAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectTradeitemAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectTradeitemAPIResponse 将 AlibabaAlihouseNewhomeProjectTradeitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectTradeitemAPIResponse(v *AlibabaAlihouseNewhomeProjectTradeitemAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectTradeitemAPIResponse.Put(v)
}
