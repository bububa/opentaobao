package fpm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSpOpenPaymentSyncAPIResponse 付款单同步 API返回值
// alibaba.sp.open.payment.sync
//
// 新康众弹外同步付款数据
type AlibabaSpOpenPaymentSyncAPIResponse struct {
	model.CommonResponse
	AlibabaSpOpenPaymentSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSpOpenPaymentSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSpOpenPaymentSyncAPIResponseModel).Reset()
}

// AlibabaSpOpenPaymentSyncAPIResponseModel is 付款单同步 成功返回结果
type AlibabaSpOpenPaymentSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_sp_open_payment_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果处理消息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 是否处理成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSpOpenPaymentSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMessage = ""
	m.IsSuccess = false
}

var poolAlibabaSpOpenPaymentSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSpOpenPaymentSyncAPIResponse)
	},
}

// GetAlibabaSpOpenPaymentSyncAPIResponse 从 sync.Pool 获取 AlibabaSpOpenPaymentSyncAPIResponse
func GetAlibabaSpOpenPaymentSyncAPIResponse() *AlibabaSpOpenPaymentSyncAPIResponse {
	return poolAlibabaSpOpenPaymentSyncAPIResponse.Get().(*AlibabaSpOpenPaymentSyncAPIResponse)
}

// ReleaseAlibabaSpOpenPaymentSyncAPIResponse 将 AlibabaSpOpenPaymentSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSpOpenPaymentSyncAPIResponse(v *AlibabaSpOpenPaymentSyncAPIResponse) {
	v.Reset()
	poolAlibabaSpOpenPaymentSyncAPIResponse.Put(v)
}
