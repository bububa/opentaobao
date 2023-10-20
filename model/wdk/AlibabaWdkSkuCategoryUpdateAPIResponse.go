package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCategoryUpdateAPIResponse 商家类目修改接口 API返回值
// alibaba.wdk.sku.category.update
//
// 商家类目修改接口
type AlibabaWdkSkuCategoryUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuCategoryUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCategoryUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuCategoryUpdateAPIResponseModel).Reset()
}

// AlibabaWdkSkuCategoryUpdateAPIResponseModel is 商家类目修改接口 成功返回结果
type AlibabaWdkSkuCategoryUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_category_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuCategoryUpdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCategoryUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuCategoryUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCategoryUpdateAPIResponse)
	},
}

// GetAlibabaWdkSkuCategoryUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuCategoryUpdateAPIResponse
func GetAlibabaWdkSkuCategoryUpdateAPIResponse() *AlibabaWdkSkuCategoryUpdateAPIResponse {
	return poolAlibabaWdkSkuCategoryUpdateAPIResponse.Get().(*AlibabaWdkSkuCategoryUpdateAPIResponse)
}

// ReleaseAlibabaWdkSkuCategoryUpdateAPIResponse 将 AlibabaWdkSkuCategoryUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuCategoryUpdateAPIResponse(v *AlibabaWdkSkuCategoryUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuCategoryUpdateAPIResponse.Put(v)
}
