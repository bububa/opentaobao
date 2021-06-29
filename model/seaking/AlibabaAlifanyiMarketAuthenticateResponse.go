package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方授权 APIResponse
alibaba.alifanyi.market.authenticate

第三方授权，获取授权码
*/
type AlibabaAlifanyiMarketAuthenticateAPIResponse struct {
    model.CommonResponse
    AlibabaAlifanyiMarketAuthenticateResponse
}

type AlibabaAlifanyiMarketAuthenticateResponse struct {
    XMLName xml.Name `xml:"alibaba_alifanyi_market_authenticate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 授权码
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // 错误描述
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`

    
    // 错误码
    
    BizErrorCode   int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`

    
    // 接口是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
