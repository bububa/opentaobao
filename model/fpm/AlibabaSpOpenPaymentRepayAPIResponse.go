package fpm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSpOpenPaymentRepayAPIResponse 智付重新打款 API返回值
// alibaba.sp.open.payment.repay
//
// 智付重新打款
type AlibabaSpOpenPaymentRepayAPIResponse struct {
	model.CommonResponse
	AlibabaSpOpenPaymentRepayAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSpOpenPaymentRepayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSpOpenPaymentRepayAPIResponseModel).Reset()
}

// AlibabaSpOpenPaymentRepayAPIResponseModel is 智付重新打款 成功返回结果
type AlibabaSpOpenPaymentRepayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_sp_open_payment_repay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果处理消息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 是否处理成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSpOpenPaymentRepayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMessage = ""
	m.IsSuccess = false
}

var poolAlibabaSpOpenPaymentRepayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSpOpenPaymentRepayAPIResponse)
	},
}

// GetAlibabaSpOpenPaymentRepayAPIResponse 从 sync.Pool 获取 AlibabaSpOpenPaymentRepayAPIResponse
func GetAlibabaSpOpenPaymentRepayAPIResponse() *AlibabaSpOpenPaymentRepayAPIResponse {
	return poolAlibabaSpOpenPaymentRepayAPIResponse.Get().(*AlibabaSpOpenPaymentRepayAPIResponse)
}

// ReleaseAlibabaSpOpenPaymentRepayAPIResponse 将 AlibabaSpOpenPaymentRepayAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSpOpenPaymentRepayAPIResponse(v *AlibabaSpOpenPaymentRepayAPIResponse) {
	v.Reset()
	poolAlibabaSpOpenPaymentRepayAPIResponse.Put(v)
}
