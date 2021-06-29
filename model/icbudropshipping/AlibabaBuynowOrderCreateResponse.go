package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴买家buynow下单接口 APIResponse
alibaba.buynow.order.create

阿里巴巴买家下单接口
*/
type AlibabaBuynowOrderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaBuynowOrderCreateResponse
}

type AlibabaBuynowOrderCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_buynow_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Response
    
    Value   *OrderCreateResponse `json:"value,omitempty" xml:"value,omitempty"`

    
}
