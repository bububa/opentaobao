package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站dropshipping 选品token 创建 APIResponse
alibaba.dropshipping.token.create

国际站dropshipping 选品token 创建，用于让买家有权限访问我们指定的 商品场馆
*/
type AlibabaDropshippingTokenCreateAPIResponse struct {
    model.CommonResponse
    AlibabaDropshippingTokenCreateResponse
}

type AlibabaDropshippingTokenCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_dropshipping_token_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // ecology_token
    
    EcologyToken   string `json:"ecology_token,omitempty" xml:"ecology_token,omitempty"`

    
}
