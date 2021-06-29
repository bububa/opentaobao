package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品api APIResponse
alibaba.alihealth.tracecodeseller.product.search

查询商品api
*/
type AlibabaAlihealthTracecodesellerProductSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerProductSearchResponse
}

type AlibabaAlihealthTracecodesellerProductSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_product_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
