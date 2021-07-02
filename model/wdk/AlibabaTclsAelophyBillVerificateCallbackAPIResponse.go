package wdk

import (
	"encoding/xml"

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

// AlibabaTclsAelophyBillVerificateCallbackAPIResponseModel is 翱象ERP核销回调 成功返回结果
type AlibabaTclsAelophyBillVerificateCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_bill_verificate_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理结果
	ApiResult *AlibabaTclsAelophyBillVerificateCallbackApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
