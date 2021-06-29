package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育三方ID绑定接口 APIResponse
alibaba.alisports.passport.account.bindthirdid

阿里体育三方ID绑定接口
*/
type AlibabaAlisportsPassportAccountBindthirdidAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsPassportAccountBindthirdidResponse
}

type AlibabaAlisportsPassportAccountBindthirdidResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_passport_account_bindthirdid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口响应码
    
    AlispCode   string `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`

    
    // 描述
    
    AlispMsg   string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`

    
}
