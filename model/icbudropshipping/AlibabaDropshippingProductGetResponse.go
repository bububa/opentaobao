package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴dropshipping 产品信息获取 APIResponse
alibaba.dropshipping.product.get

阿里巴巴dropshipping 产品信息获取
*/
type AlibabaDropshippingProductGetAPIResponse struct {
    model.CommonResponse
    AlibabaDropshippingProductGetResponse
}

type AlibabaDropshippingProductGetResponse struct {
    XMLName xml.Name `xml:"alibaba_dropshipping_product_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // product pojo
    
    Value   []DistributionSaleProduct `json:"value,omitempty" xml:"value>distribution_sale_product,omitempty"`
    
    
}
