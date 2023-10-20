package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerproductattrsearchAPIResponse 根据商品id获取商品属性 API返回值
// alibaba.alihealth.tracecodeseller.product.attr.search
//
// 根据商品id获取商品属性
type AlibabaalihealthtracecodesellerproductattrsearchAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtracecodesellerproductattrsearchAPIResponseModel
}

// AlibabaalihealthtracecodesellerproductattrsearchAPIResponseModel is 根据商品id获取商品属性 成功返回结果
type AlibabaalihealthtracecodesellerproductattrsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_product_attr_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
