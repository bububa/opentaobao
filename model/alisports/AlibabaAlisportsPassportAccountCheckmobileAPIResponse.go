package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统--手机号验证接口 API返回值 
alibaba.alisports.passport.account.checkmobile

验证三方用户的手机号，获取对应的阿里体育会员id
*/
type AlibabaAlisportsPassportAccountCheckmobileAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsPassportAccountCheckmobileAPIResponseModel
}

// 阿里体育会员系统--手机号验证接口 成功返回结果
type AlibabaAlisportsPassportAccountCheckmobileAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alisports_passport_account_checkmobile_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alisp_msg
    AlispMsg   string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
    // alisp_code
    AlispCode   int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
    // alisp_data
    AlispData   *AlispData `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
}
