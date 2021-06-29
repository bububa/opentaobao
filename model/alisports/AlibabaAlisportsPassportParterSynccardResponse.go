package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育-卡信息同步接口 API返回值 
alibaba.alisports.passport.parter.synccard

运享通修改卡号的时候，通知更新到阿里体育和支付宝卡包中
*/
type AlibabaAlisportsPassportParterSynccardAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsPassportParterSynccardResponse
}

// 阿里体育-卡信息同步接口 成功返回结果
type AlibabaAlisportsPassportParterSynccardResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_passport_parter_synccard_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 200标识成功，其他的code为失败
    AlispCode   int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
    // 正确或错误的信息
    AlispMsg   string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
    // 返回数据
    AlispData   string `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
}
