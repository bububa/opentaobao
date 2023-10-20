package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCategoryAddAPIResponse 商家类目新增接口 API返回值
// alibaba.wdk.sku.category.add
//
// 商家类目新增接口
type AlibabaWdkSkuCategoryAddAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuCategoryAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCategoryAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuCategoryAddAPIResponseModel).Reset()
}

// AlibabaWdkSkuCategoryAddAPIResponseModel is 商家类目新增接口 成功返回结果
type AlibabaWdkSkuCategoryAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_category_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuCategoryAddApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCategoryAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuCategoryAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCategoryAddAPIResponse)
	},
}

// GetAlibabaWdkSkuCategoryAddAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuCategoryAddAPIResponse
func GetAlibabaWdkSkuCategoryAddAPIResponse() *AlibabaWdkSkuCategoryAddAPIResponse {
	return poolAlibabaWdkSkuCategoryAddAPIResponse.Get().(*AlibabaWdkSkuCategoryAddAPIResponse)
}

// ReleaseAlibabaWdkSkuCategoryAddAPIResponse 将 AlibabaWdkSkuCategoryAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuCategoryAddAPIResponse(v *AlibabaWdkSkuCategoryAddAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuCategoryAddAPIResponse.Put(v)
}
