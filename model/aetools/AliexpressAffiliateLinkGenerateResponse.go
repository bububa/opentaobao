package aetools

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟推广链接生成 API返回值 
aliexpress.affiliate.link.generate

AE联盟推广链接生成接口
*/
type AliexpressAffiliateLinkGenerateAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateLinkGenerateResponse
}

// 联盟推广链接生成 成功返回结果
type AliexpressAffiliateLinkGenerateResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_link_generate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}