package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyBillVerificateCallbackAPIResponse 翱象ERP核销回调 API返回值
// alibaba.tcls.aelophy.bill.verificate.callback
//
// 翱象ERP核销回调
type AlibabaTclsAelophyBillVerificateCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyBillVerificateCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyBillVerificateCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyBillVerificateCallbackAPIResponseModel).Reset()
}

// AlibabaTclsAelophyBillVerificateCallbackAPIResponseModel is 翱象ERP核销回调 成功返回结果
type AlibabaTclsAelophyBillVerificateCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_bill_verificate_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理结果
	ApiResult *AlibabaTclsAelophyBillVerificateCallbackApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyBillVerificateCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyBillVerificateCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyBillVerificateCallbackAPIResponse)
	},
}

// GetAlibabaTclsAelophyBillVerificateCallbackAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyBillVerificateCallbackAPIResponse
func GetAlibabaTclsAelophyBillVerificateCallbackAPIResponse() *AlibabaTclsAelophyBillVerificateCallbackAPIResponse {
	return poolAlibabaTclsAelophyBillVerificateCallbackAPIResponse.Get().(*AlibabaTclsAelophyBillVerificateCallbackAPIResponse)
}

// ReleaseAlibabaTclsAelophyBillVerificateCallbackAPIResponse 将 AlibabaTclsAelophyBillVerificateCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyBillVerificateCallbackAPIResponse(v *AlibabaTclsAelophyBillVerificateCallbackAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyBillVerificateCallbackAPIResponse.Put(v)
}
