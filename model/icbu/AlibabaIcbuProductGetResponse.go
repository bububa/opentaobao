package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获得单个商品详情 APIResponse
alibaba.icbu.product.get

获取商品详情
*/
type AlibabaIcbuProductGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuProductGetResponse
}

type AlibabaIcbuProductGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_product_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 单个商品详情
    
    Product   *AlibabaProductResponse `json:"product,omitempty" xml:"product,omitempty"`

    
}
