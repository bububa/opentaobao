package aetools

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟推广链接生成 APIResponse
aliexpress.affiliate.link.generate

AE联盟推广链接生成接口
*/
type AliexpressAffiliateLinkGenerateAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateLinkGenerateResponse
}

type AliexpressAffiliateLinkGenerateResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_link_generate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`

    
}
