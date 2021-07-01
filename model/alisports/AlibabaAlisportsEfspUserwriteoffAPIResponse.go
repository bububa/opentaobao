package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户核销 API返回值 
alibaba.alisports.efsp.userwriteoff

用户核销
*/
type AlibabaAlisportsEfspUserwriteoffAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsEfspUserwriteoffAPIResponseModel
}

// 用户核销 成功返回结果
type AlibabaAlisportsEfspUserwriteoffAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alisports_efsp_userwriteoff_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    Result   *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`
}
