package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品id获取商品属性 APIResponse
alibaba.alihealth.tracecodeseller.product.attr.search

根据商品id获取商品属性
*/
type AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerProductAttrSearchResponse
}

type AlibabaAlihealthTracecodesellerProductAttrSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_product_attr_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
