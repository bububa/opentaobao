package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsMessageDirectBatchsendAPIResponse OAID批量发送，支持明文手机号发送 API返回值
// taobao.jst.sms.message.direct.batchsend
//
// OAID批量发送，支持明文手机号发送
type TaobaoJstSmsMessageDirectBatchsendAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsMessageDirectBatchsendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstSmsMessageDirectBatchsendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsMessageDirectBatchsendAPIResponseModel).Reset()
}

// TaobaoJstSmsMessageDirectBatchsendAPIResponseModel is OAID批量发送，支持明文手机号发送 成功返回结果
type TaobaoJstSmsMessageDirectBatchsendAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_message_direct_batchsend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 短信回执码
	Module string `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstSmsMessageDirectBatchsendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = ""
}

var poolTaobaoJstSmsMessageDirectBatchsendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsMessageDirectBatchsendAPIResponse)
	},
}

// GetTaobaoJstSmsMessageDirectBatchsendAPIResponse 从 sync.Pool 获取 TaobaoJstSmsMessageDirectBatchsendAPIResponse
func GetTaobaoJstSmsMessageDirectBatchsendAPIResponse() *TaobaoJstSmsMessageDirectBatchsendAPIResponse {
	return poolTaobaoJstSmsMessageDirectBatchsendAPIResponse.Get().(*TaobaoJstSmsMessageDirectBatchsendAPIResponse)
}

// ReleaseTaobaoJstSmsMessageDirectBatchsendAPIResponse 将 TaobaoJstSmsMessageDirectBatchsendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsMessageDirectBatchsendAPIResponse(v *TaobaoJstSmsMessageDirectBatchsendAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsMessageDirectBatchsendAPIResponse.Put(v)
}
