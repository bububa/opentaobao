package aesolution

import (
	"encoding/xml"

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
