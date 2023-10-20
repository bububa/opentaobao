package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemCategoryUpdateAPIResponse 修改类目 API返回值
// alibaba.wdk.item.category.update
//
// 修改类目
type AlibabaWdkItemCategoryUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemCategoryUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemCategoryUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemCategoryUpdateAPIResponseModel).Reset()
}

// AlibabaWdkItemCategoryUpdateAPIResponseModel is 修改类目 成功返回结果
type AlibabaWdkItemCategoryUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_category_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemCategoryUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemCategoryUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemCategoryUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemCategoryUpdateAPIResponse)
	},
}

// GetAlibabaWdkItemCategoryUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkItemCategoryUpdateAPIResponse
func GetAlibabaWdkItemCategoryUpdateAPIResponse() *AlibabaWdkItemCategoryUpdateAPIResponse {
	return poolAlibabaWdkItemCategoryUpdateAPIResponse.Get().(*AlibabaWdkItemCategoryUpdateAPIResponse)
}

// ReleaseAlibabaWdkItemCategoryUpdateAPIResponse 将 AlibabaWdkItemCategoryUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemCategoryUpdateAPIResponse(v *AlibabaWdkItemCategoryUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemCategoryUpdateAPIResponse.Put(v)
}
