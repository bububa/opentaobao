package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeitemRelationAPIResponse 货独立绑定货品 API返回值
// alibaba.alihouse.newhome.tradeitem.relation
//
// 货独立绑定货品
type AlibabaAlihouseNewhomeTradeitemRelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeTradeitemRelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeTradeitemRelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeTradeitemRelationAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeTradeitemRelationAPIResponseModel is 货独立绑定货品 成功返回结果
type AlibabaAlihouseNewhomeTradeitemRelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_tradeitem_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeTradeitemRelationResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeTradeitemRelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeTradeitemRelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeTradeitemRelationAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeTradeitemRelationAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeTradeitemRelationAPIResponse
func GetAlibabaAlihouseNewhomeTradeitemRelationAPIResponse() *AlibabaAlihouseNewhomeTradeitemRelationAPIResponse {
	return poolAlibabaAlihouseNewhomeTradeitemRelationAPIResponse.Get().(*AlibabaAlihouseNewhomeTradeitemRelationAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeTradeitemRelationAPIResponse 将 AlibabaAlihouseNewhomeTradeitemRelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeTradeitemRelationAPIResponse(v *AlibabaAlihouseNewhomeTradeitemRelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeTradeitemRelationAPIResponse.Put(v)
}
