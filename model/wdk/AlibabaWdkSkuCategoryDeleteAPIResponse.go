package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCategoryDeleteAPIResponse 商家类目删除接口 API返回值
// alibaba.wdk.sku.category.delete
//
// 商家类目删除接口
type AlibabaWdkSkuCategoryDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuCategoryDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCategoryDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuCategoryDeleteAPIResponseModel).Reset()
}

// AlibabaWdkSkuCategoryDeleteAPIResponseModel is 商家类目删除接口 成功返回结果
type AlibabaWdkSkuCategoryDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_category_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuCategoryDeleteApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCategoryDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuCategoryDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCategoryDeleteAPIResponse)
	},
}

// GetAlibabaWdkSkuCategoryDeleteAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuCategoryDeleteAPIResponse
func GetAlibabaWdkSkuCategoryDeleteAPIResponse() *AlibabaWdkSkuCategoryDeleteAPIResponse {
	return poolAlibabaWdkSkuCategoryDeleteAPIResponse.Get().(*AlibabaWdkSkuCategoryDeleteAPIResponse)
}

// ReleaseAlibabaWdkSkuCategoryDeleteAPIResponse 将 AlibabaWdkSkuCategoryDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuCategoryDeleteAPIResponse(v *AlibabaWdkSkuCategoryDeleteAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuCategoryDeleteAPIResponse.Put(v)
}
