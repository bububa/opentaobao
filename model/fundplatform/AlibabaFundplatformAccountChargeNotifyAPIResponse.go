package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountChargeNotifyAPIResponse 账户充值成功通知 API返回值
// alibaba.fundplatform.account.charge.notify
//
// 通知外部业务方充值成功
type AlibabaFundplatformAccountChargeNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformAccountChargeNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFundplatformAccountChargeNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFundplatformAccountChargeNotifyAPIResponseModel).Reset()
}

// AlibabaFundplatformAccountChargeNotifyAPIResponseModel is 账户充值成功通知 成功返回结果
type AlibabaFundplatformAccountChargeNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_account_charge_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理消息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 处理错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 处理结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFundplatformAccountChargeNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMessage = ""
	m.ResultCode = ""
	m.Success = false
}

var poolAlibabaFundplatformAccountChargeNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformAccountChargeNotifyAPIResponse)
	},
}

// GetAlibabaFundplatformAccountChargeNotifyAPIResponse 从 sync.Pool 获取 AlibabaFundplatformAccountChargeNotifyAPIResponse
func GetAlibabaFundplatformAccountChargeNotifyAPIResponse() *AlibabaFundplatformAccountChargeNotifyAPIResponse {
	return poolAlibabaFundplatformAccountChargeNotifyAPIResponse.Get().(*AlibabaFundplatformAccountChargeNotifyAPIResponse)
}

// ReleaseAlibabaFundplatformAccountChargeNotifyAPIResponse 将 AlibabaFundplatformAccountChargeNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFundplatformAccountChargeNotifyAPIResponse(v *AlibabaFundplatformAccountChargeNotifyAPIResponse) {
	v.Reset()
	poolAlibabaFundplatformAccountChargeNotifyAPIResponse.Put(v)
}
