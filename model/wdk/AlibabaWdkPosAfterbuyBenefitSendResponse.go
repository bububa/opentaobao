package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
生态pos购后发放权益 API返回值 
alibaba.wdk.pos.afterbuy.benefit.send

生态pos购后发放权益接口开放
*/
type AlibabaWdkPosAfterbuyBenefitSendAPIResponse struct {
    model.CommonResponse
    AlibabaWdkPosAfterbuyBenefitSendResponse
}

// 生态pos购后发放权益 成功返回结果
type AlibabaWdkPosAfterbuyBenefitSendResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_pos_afterbuy_benefit_send_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}
