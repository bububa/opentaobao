package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionSellerCategoryTreeQueryAPIResponse aliexpress.solution.seller.category.tree.query API返回值
// aliexpress.solution.seller.category.tree.query
//
// API for seller to query the category tree. Support only displaying the categories which seller have permissions to publish products.
type AliexpressSolutionSellerCategoryTreeQueryAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionSellerCategoryTreeQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionSellerCategoryTreeQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionSellerCategoryTreeQueryAPIResponseModel).Reset()
}

// AliexpressSolutionSellerCategoryTreeQueryAPIResponseModel is aliexpress.solution.seller.category.tree.query 成功返回结果
type AliexpressSolutionSellerCategoryTreeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_seller_category_tree_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// children category list under category_id
	ChildrenCategoryList []CategoryInfo `json:"children_category_list,omitempty" xml:"children_category_list>category_info,omitempty"`
	// whether success or not
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionSellerCategoryTreeQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ChildrenCategoryList = m.ChildrenCategoryList[:0]
	m.IsSuccess = false
}

var poolAliexpressSolutionSellerCategoryTreeQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionSellerCategoryTreeQueryAPIResponse)
	},
}

// GetAliexpressSolutionSellerCategoryTreeQueryAPIResponse 从 sync.Pool 获取 AliexpressSolutionSellerCategoryTreeQueryAPIResponse
func GetAliexpressSolutionSellerCategoryTreeQueryAPIResponse() *AliexpressSolutionSellerCategoryTreeQueryAPIResponse {
	return poolAliexpressSolutionSellerCategoryTreeQueryAPIResponse.Get().(*AliexpressSolutionSellerCategoryTreeQueryAPIResponse)
}

// ReleaseAliexpressSolutionSellerCategoryTreeQueryAPIResponse 将 AliexpressSolutionSellerCategoryTreeQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionSellerCategoryTreeQueryAPIResponse(v *AliexpressSolutionSellerCategoryTreeQueryAPIResponse) {
	v.Reset()
	poolAliexpressSolutionSellerCategoryTreeQueryAPIResponse.Put(v)
}
