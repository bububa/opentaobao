package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductCategorySuggestAPIResponse Suggest product categories API返回值
// aliexpress.solution.product.category.suggest
//
// Suggest product categories by title and image.
type AliexpressSolutionProductCategorySuggestAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionProductCategorySuggestAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionProductCategorySuggestAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionProductCategorySuggestAPIResponseModel).Reset()
}

// AliexpressSolutionProductCategorySuggestAPIResponseModel is Suggest product categories 成功返回结果
type AliexpressSolutionProductCategorySuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_product_category_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// suggest category list, order by priority
	SuggestCategoryList []SuggestCategory `json:"suggest_category_list,omitempty" xml:"suggest_category_list>suggest_category,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionProductCategorySuggestAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SuggestCategoryList = m.SuggestCategoryList[:0]
}

var poolAliexpressSolutionProductCategorySuggestAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionProductCategorySuggestAPIResponse)
	},
}

// GetAliexpressSolutionProductCategorySuggestAPIResponse 从 sync.Pool 获取 AliexpressSolutionProductCategorySuggestAPIResponse
func GetAliexpressSolutionProductCategorySuggestAPIResponse() *AliexpressSolutionProductCategorySuggestAPIResponse {
	return poolAliexpressSolutionProductCategorySuggestAPIResponse.Get().(*AliexpressSolutionProductCategorySuggestAPIResponse)
}

// ReleaseAliexpressSolutionProductCategorySuggestAPIResponse 将 AliexpressSolutionProductCategorySuggestAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionProductCategorySuggestAPIResponse(v *AliexpressSolutionProductCategorySuggestAPIResponse) {
	v.Reset()
	poolAliexpressSolutionProductCategorySuggestAPIResponse.Put(v)
}
