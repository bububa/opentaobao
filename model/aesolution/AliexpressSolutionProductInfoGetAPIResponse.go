package aesolution

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductInfoGetAPIResponse Get Single Product Info API返回值
// aliexpress.solution.product.info.get
//
// Get Single Product Info
type AliexpressSolutionProductInfoGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionProductInfoGetAPIResponseModel
}

// AliexpressSolutionProductInfoGetAPIResponseModel is Get Single Product Info 成功返回结果
type AliexpressSolutionProductInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_product_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *GlobalAeopFindProductResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
