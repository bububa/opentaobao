package aesolution

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionproductcategorysuggestAPIResponse Suggest product categories API返回值
// aliexpress.solution.product.category.suggest
//
// Suggest product categories by title and image.
type AliexpresssolutionproductcategorysuggestAPIResponse struct {
	model.CommonResponse
	AliexpresssolutionproductcategorysuggestAPIResponseModel
}

// AliexpresssolutionproductcategorysuggestAPIResponseModel is Suggest product categories 成功返回结果
type AliexpresssolutionproductcategorysuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_product_category_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// suggest category list, order by priority
	SuggestCategoryList []SuggestCategory `json:"suggest_category_list,omitempty" xml:"suggest_category_list>suggest_category,omitempty"`
}
