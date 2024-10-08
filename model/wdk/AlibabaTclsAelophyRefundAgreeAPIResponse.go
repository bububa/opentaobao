package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundAgreeAPIResponse saas 售后逆向 商户同意用户逆向申请 API返回值
// alibaba.tcls.aelophy.refund.agree
//
// saas 售后逆向 商户同意用户逆向申请
type AlibabaTclsAelophyRefundAgreeAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyRefundAgreeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyRefundAgreeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyRefundAgreeAPIResponseModel).Reset()
}

// AlibabaTclsAelophyRefundAgreeAPIResponseModel is saas 售后逆向 商户同意用户逆向申请 成功返回结果
type AlibabaTclsAelophyRefundAgreeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabaTclsAelophyRefundAgreeApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyRefundAgreeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTclsAelophyRefundAgreeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundAgreeAPIResponse)
	},
}

// GetAlibabaTclsAelophyRefundAgreeAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyRefundAgreeAPIResponse
func GetAlibabaTclsAelophyRefundAgreeAPIResponse() *AlibabaTclsAelophyRefundAgreeAPIResponse {
	return poolAlibabaTclsAelophyRefundAgreeAPIResponse.Get().(*AlibabaTclsAelophyRefundAgreeAPIResponse)
}

// ReleaseAlibabaTclsAelophyRefundAgreeAPIResponse 将 AlibabaTclsAelophyRefundAgreeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyRefundAgreeAPIResponse(v *AlibabaTclsAelophyRefundAgreeAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyRefundAgreeAPIResponse.Put(v)
}
