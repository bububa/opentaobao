package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴商品运费计算查询接口 APIResponse
alibaba.shipping.freight.calculate

阿里巴巴商品运费计算查询接口
*/
type AlibabaShippingFreightCalculateAPIResponse struct {
    model.CommonResponse
    AlibabaShippingFreightCalculateResponse
}

type AlibabaShippingFreightCalculateResponse struct {
    XMLName xml.Name `xml:"alibaba_shipping_freight_calculate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // pojo
    
    Values   []Value `json:"values,omitempty" xml:"values>value,omitempty"`
    
    
}
