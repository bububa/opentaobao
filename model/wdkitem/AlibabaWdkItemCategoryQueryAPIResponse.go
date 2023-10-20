package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemCategoryQueryAPIResponse 类目查询接口 API返回值
// alibaba.wdk.item.category.query
//
// 类目查询接口
type AlibabaWdkItemCategoryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemCategoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemCategoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemCategoryQueryAPIResponseModel).Reset()
}

// AlibabaWdkItemCategoryQueryAPIResponseModel is 类目查询接口 成功返回结果
type AlibabaWdkItemCategoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_category_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemCategoryQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemCategoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemCategoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemCategoryQueryAPIResponse)
	},
}

// GetAlibabaWdkItemCategoryQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkItemCategoryQueryAPIResponse
func GetAlibabaWdkItemCategoryQueryAPIResponse() *AlibabaWdkItemCategoryQueryAPIResponse {
	return poolAlibabaWdkItemCategoryQueryAPIResponse.Get().(*AlibabaWdkItemCategoryQueryAPIResponse)
}

// ReleaseAlibabaWdkItemCategoryQueryAPIResponse 将 AlibabaWdkItemCategoryQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemCategoryQueryAPIResponse(v *AlibabaWdkItemCategoryQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemCategoryQueryAPIResponse.Put(v)
}
