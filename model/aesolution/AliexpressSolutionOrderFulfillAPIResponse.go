package aesolution

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionorderfulfillAPIResponse fulfill order API返回值
// aliexpress.solution.order.fulfill
//
// fulfill order for seller
type AliexpresssolutionorderfulfillAPIResponse struct {
	model.CommonResponse
	AliexpresssolutionorderfulfillAPIResponseModel
}

// AliexpresssolutionorderfulfillAPIResponseModel is fulfill order 成功返回结果
type AliexpresssolutionorderfulfillAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_order_fulfill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
