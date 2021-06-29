package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户取消订单 API返回值 
alibaba.alisports.efsp.usercancelorder

用户取消订单
*/
type AlibabaAlisportsEfspUsercancelorderAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsEfspUsercancelorderResponse
}

// 用户取消订单 成功返回结果
type AlibabaAlisportsEfspUsercancelorderResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_efsp_usercancelorder_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    Result   *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`
}
