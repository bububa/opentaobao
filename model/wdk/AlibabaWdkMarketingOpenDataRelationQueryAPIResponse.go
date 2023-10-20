package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenDataRelationQueryAPIResponse 数据关联关系查询 API返回值
// alibaba.wdk.marketing.open.data.relation.query
//
// 数据关联关系查询
type AlibabaWdkMarketingOpenDataRelationQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingOpenDataRelationQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenDataRelationQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingOpenDataRelationQueryAPIResponseModel).Reset()
}

// AlibabaWdkMarketingOpenDataRelationQueryAPIResponseModel is 数据关联关系查询 成功返回结果
type AlibabaWdkMarketingOpenDataRelationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_data_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenDataRelationQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingOpenDataRelationQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingOpenDataRelationQueryAPIResponse)
	},
}

// GetAlibabaWdkMarketingOpenDataRelationQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingOpenDataRelationQueryAPIResponse
func GetAlibabaWdkMarketingOpenDataRelationQueryAPIResponse() *AlibabaWdkMarketingOpenDataRelationQueryAPIResponse {
	return poolAlibabaWdkMarketingOpenDataRelationQueryAPIResponse.Get().(*AlibabaWdkMarketingOpenDataRelationQueryAPIResponse)
}

// ReleaseAlibabaWdkMarketingOpenDataRelationQueryAPIResponse 将 AlibabaWdkMarketingOpenDataRelationQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingOpenDataRelationQueryAPIResponse(v *AlibabaWdkMarketingOpenDataRelationQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingOpenDataRelationQueryAPIResponse.Put(v)
}
