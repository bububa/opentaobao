package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户同意用户逆向申请 API返回值 
alibaba.tcls.aelophy.refund.agree

saas 售后逆向 商户同意用户逆向申请
*/
type AlibabaTclsAelophyRefundAgreeAPIResponse struct {
    model.CommonResponse
    AlibabaTclsAelophyRefundAgreeAPIResponseModel
}

// saas 售后逆向 商户同意用户逆向申请 成功返回结果
type AlibabaTclsAelophyRefundAgreeAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_agree_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Result   *AlibabaTclsAelophyRefundAgreeApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
