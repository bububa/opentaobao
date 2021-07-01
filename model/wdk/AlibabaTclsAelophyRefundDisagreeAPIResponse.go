package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyRefundDisagreeAPIResponse
saas 售后逆向 商户拒绝用户逆向申请 API返回值
alibaba.tcls.aelophy.refund.disagree

saas 售后逆向 商户拒绝用户逆向申请 */
type AlibabaTclsAelophyRefundDisagreeAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyRefundDisagreeAPIResponseModel
}

// AlibabaTclsAelophyRefundDisagreeAPIResponseModel is saas 售后逆向 商户拒绝用户逆向申请 成功返回结果
type AlibabaTclsAelophyRefundDisagreeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_disagree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabaTclsAelophyRefundDisagreeApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
