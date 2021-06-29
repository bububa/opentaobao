package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计算补助金额 APIResponse
alibaba.alisports.efsp.countsubsidy

计算补助金额
*/
type AlibabaAlisportsEfspCountsubsidyAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsEfspCountsubsidyResponse
}

type AlibabaAlisportsEfspCountsubsidyResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_efsp_countsubsidy_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
