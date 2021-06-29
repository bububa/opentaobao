package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统--取消三方关联接口 APIResponse
alibaba.alisports.passport.account.delrelation

阿里体育会员系统--取消三方关联接口
*/
type AlibabaAlisportsPassportAccountDelrelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsPassportAccountDelrelationResponse
}

type AlibabaAlisportsPassportAccountDelrelationResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_passport_account_delrelation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alisp_code
    
    AlispCode   int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`

    
    // alisp_msg
    
    AlispMsg   string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`

    
    // alisp_data
    
    AlispData   string `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`

    
}
