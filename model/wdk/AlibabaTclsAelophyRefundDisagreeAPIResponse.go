package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundDisagreeAPIResponse saas 售后逆向 商户拒绝用户逆向申请 API返回值
// alibaba.tcls.aelophy.refund.disagree
//
// saas 售后逆向 商户拒绝用户逆向申请
type AlibabaTclsAelophyRefundDisagreeAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyRefundDisagreeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyRefundDisagreeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyRefundDisagreeAPIResponseModel).Reset()
}

// AlibabaTclsAelophyRefundDisagreeAPIResponseModel is saas 售后逆向 商户拒绝用户逆向申请 成功返回结果
type AlibabaTclsAelophyRefundDisagreeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_disagree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabaTclsAelophyRefundDisagreeApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyRefundDisagreeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTclsAelophyRefundDisagreeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundDisagreeAPIResponse)
	},
}

// GetAlibabaTclsAelophyRefundDisagreeAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyRefundDisagreeAPIResponse
func GetAlibabaTclsAelophyRefundDisagreeAPIResponse() *AlibabaTclsAelophyRefundDisagreeAPIResponse {
	return poolAlibabaTclsAelophyRefundDisagreeAPIResponse.Get().(*AlibabaTclsAelophyRefundDisagreeAPIResponse)
}

// ReleaseAlibabaTclsAelophyRefundDisagreeAPIResponse 将 AlibabaTclsAelophyRefundDisagreeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyRefundDisagreeAPIResponse(v *AlibabaTclsAelophyRefundDisagreeAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyRefundDisagreeAPIResponse.Put(v)
}
