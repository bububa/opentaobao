package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinTaSmsNumSendAPIResponse 短信发送 API返回值
// alibaba.aliqin.ta.sms.num.send
//
// 短信发送
type AlibabaAliqinTaSmsNumSendAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinTaSmsNumSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinTaSmsNumSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinTaSmsNumSendAPIResponseModel).Reset()
}

// AlibabaAliqinTaSmsNumSendAPIResponseModel is 短信发送 成功返回结果
type AlibabaAliqinTaSmsNumSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_ta_sms_num_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAliqinTaSmsNumSendBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinTaSmsNumSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinTaSmsNumSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTaSmsNumSendAPIResponse)
	},
}

// GetAlibabaAliqinTaSmsNumSendAPIResponse 从 sync.Pool 获取 AlibabaAliqinTaSmsNumSendAPIResponse
func GetAlibabaAliqinTaSmsNumSendAPIResponse() *AlibabaAliqinTaSmsNumSendAPIResponse {
	return poolAlibabaAliqinTaSmsNumSendAPIResponse.Get().(*AlibabaAliqinTaSmsNumSendAPIResponse)
}

// ReleaseAlibabaAliqinTaSmsNumSendAPIResponse 将 AlibabaAliqinTaSmsNumSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinTaSmsNumSendAPIResponse(v *AlibabaAliqinTaSmsNumSendAPIResponse) {
	v.Reset()
	poolAlibabaAliqinTaSmsNumSendAPIResponse.Put(v)
}
