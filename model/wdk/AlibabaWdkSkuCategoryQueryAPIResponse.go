package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCategoryQueryAPIResponse 商家类目获取接口 API返回值
// alibaba.wdk.sku.category.query
//
// 商家类目获取接口
type AlibabaWdkSkuCategoryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuCategoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCategoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuCategoryQueryAPIResponseModel).Reset()
}

// AlibabaWdkSkuCategoryQueryAPIResponseModel is 商家类目获取接口 成功返回结果
type AlibabaWdkSkuCategoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_category_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuCategoryQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCategoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuCategoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCategoryQueryAPIResponse)
	},
}

// GetAlibabaWdkSkuCategoryQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuCategoryQueryAPIResponse
func GetAlibabaWdkSkuCategoryQueryAPIResponse() *AlibabaWdkSkuCategoryQueryAPIResponse {
	return poolAlibabaWdkSkuCategoryQueryAPIResponse.Get().(*AlibabaWdkSkuCategoryQueryAPIResponse)
}

// ReleaseAlibabaWdkSkuCategoryQueryAPIResponse 将 AlibabaWdkSkuCategoryQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuCategoryQueryAPIResponse(v *AlibabaWdkSkuCategoryQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuCategoryQueryAPIResponse.Put(v)
}
