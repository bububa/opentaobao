package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundCsapplyrenderAPIResponse 商家代客售后逆向申请渲染获取 API返回值
// alibaba.tcls.aelophy.refund.csapplyrender
//
// 提供商家代客售后逆向申请渲染获取的接口
type AlibabaTclsAelophyRefundCsapplyrenderAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyRefundCsapplyrenderAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyRefundCsapplyrenderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyRefundCsapplyrenderAPIResponseModel).Reset()
}

// AlibabaTclsAelophyRefundCsapplyrenderAPIResponseModel is 商家代客售后逆向申请渲染获取 成功返回结果
type AlibabaTclsAelophyRefundCsapplyrenderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_csapplyrender_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	ApiResult *AlibabaTclsAelophyRefundCsapplyrenderApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyRefundCsapplyrenderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyRefundCsapplyrenderAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundCsapplyrenderAPIResponse)
	},
}

// GetAlibabaTclsAelophyRefundCsapplyrenderAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyRefundCsapplyrenderAPIResponse
func GetAlibabaTclsAelophyRefundCsapplyrenderAPIResponse() *AlibabaTclsAelophyRefundCsapplyrenderAPIResponse {
	return poolAlibabaTclsAelophyRefundCsapplyrenderAPIResponse.Get().(*AlibabaTclsAelophyRefundCsapplyrenderAPIResponse)
}

// ReleaseAlibabaTclsAelophyRefundCsapplyrenderAPIResponse 将 AlibabaTclsAelophyRefundCsapplyrenderAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyRefundCsapplyrenderAPIResponse(v *AlibabaTclsAelophyRefundCsapplyrenderAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyRefundCsapplyrenderAPIResponse.Put(v)
}
