package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
重发核销码 API返回值 
tmall.servicecenter.workcard.verifycode.resend

重发核销码
*/
type TmallServicecenterWorkcardVerifycodeResendAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardVerifycodeResendResponse
}

// 重发核销码 成功返回结果
type TmallServicecenterWorkcardVerifycodeResendResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_verifycode_resend_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果
    Result   *TmallServicecenterWorkcardVerifycodeResendResult `json:"result,omitempty" xml:"result,omitempty"`
}
