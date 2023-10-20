package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemBrandQueryAPIResponse 品牌信息查询 API返回值
// alibaba.wdk.item.brand.query
//
// 品牌信息查询
type AlibabaWdkItemBrandQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemBrandQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemBrandQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemBrandQueryAPIResponseModel).Reset()
}

// AlibabaWdkItemBrandQueryAPIResponseModel is 品牌信息查询 成功返回结果
type AlibabaWdkItemBrandQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_brand_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemBrandQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemBrandQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemBrandQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemBrandQueryAPIResponse)
	},
}

// GetAlibabaWdkItemBrandQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkItemBrandQueryAPIResponse
func GetAlibabaWdkItemBrandQueryAPIResponse() *AlibabaWdkItemBrandQueryAPIResponse {
	return poolAlibabaWdkItemBrandQueryAPIResponse.Get().(*AlibabaWdkItemBrandQueryAPIResponse)
}

// ReleaseAlibabaWdkItemBrandQueryAPIResponse 将 AlibabaWdkItemBrandQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemBrandQueryAPIResponse(v *AlibabaWdkItemBrandQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemBrandQueryAPIResponse.Put(v)
}
