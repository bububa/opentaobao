package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerproductsearchAPIResponse 查询商品api API返回值
// alibaba.alihealth.tracecodeseller.product.search
//
// 查询商品api
type AlibabaalihealthtracecodesellerproductsearchAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtracecodesellerproductsearchAPIResponseModel
}

// AlibabaalihealthtracecodesellerproductsearchAPIResponseModel is 查询商品api 成功返回结果
type AlibabaalihealthtracecodesellerproductsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_product_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
