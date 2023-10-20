package aesolution

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionordergetAPIResponse get order list API返回值
// aliexpress.solution.order.get
//
// Get Order List from AliExpress
type AliexpresssolutionordergetAPIResponse struct {
	model.CommonResponse
	AliexpresssolutionordergetAPIResponseModel
}

// AliexpresssolutionordergetAPIResponseModel is get order list 成功返回结果
type AliexpresssolutionordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PaginationResult `json:"result,omitempty" xml:"result,omitempty"`
}
