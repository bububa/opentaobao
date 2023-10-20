package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefundagreeAPIResponse saas 售后逆向 商户同意用户逆向申请 API返回值
// alibaba.tcls.aelophy.refund.agree
//
// saas 售后逆向 商户同意用户逆向申请
type AlibabatclsaelophyrefundagreeAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophyrefundagreeAPIResponseModel
}

// AlibabatclsaelophyrefundagreeAPIResponseModel is saas 售后逆向 商户同意用户逆向申请 成功返回结果
type AlibabatclsaelophyrefundagreeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabatclsaelophyrefundagreeApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
