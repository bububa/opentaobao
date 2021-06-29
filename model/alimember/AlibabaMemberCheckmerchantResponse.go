package alimember

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验商家身份 APIResponse
alibaba.member.checkmerchant

校验商家身份
*/
type AlibabaMemberCheckmerchantAPIResponse struct {
    model.CommonResponse
    AlibabaMemberCheckmerchantResponse
}

type AlibabaMemberCheckmerchantResponse struct {
    XMLName xml.Name `xml:"alibaba_member_checkmerchant_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // message
    
    ReturnMessage   string `json:"return_message,omitempty" xml:"return_message,omitempty"`

    
    // code
    
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`

    
}
