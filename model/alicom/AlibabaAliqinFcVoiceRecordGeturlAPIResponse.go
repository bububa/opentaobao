package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcVoiceRecordGeturlAPIResponse 录音文件下载 API返回值
// alibaba.aliqin.fc.voice.record.geturl
//
// 完成录音文件的下载地址获取操作
type AlibabaAliqinFcVoiceRecordGeturlAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcVoiceRecordGeturlAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcVoiceRecordGeturlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcVoiceRecordGeturlAPIResponseModel).Reset()
}

// AlibabaAliqinFcVoiceRecordGeturlAPIResponseModel is 录音文件下载 成功返回结果
type AlibabaAliqinFcVoiceRecordGeturlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_record_geturl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcVoiceRecordGeturlResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcVoiceRecordGeturlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcVoiceRecordGeturlAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcVoiceRecordGeturlAPIResponse)
	},
}

// GetAlibabaAliqinFcVoiceRecordGeturlAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcVoiceRecordGeturlAPIResponse
func GetAlibabaAliqinFcVoiceRecordGeturlAPIResponse() *AlibabaAliqinFcVoiceRecordGeturlAPIResponse {
	return poolAlibabaAliqinFcVoiceRecordGeturlAPIResponse.Get().(*AlibabaAliqinFcVoiceRecordGeturlAPIResponse)
}

// ReleaseAlibabaAliqinFcVoiceRecordGeturlAPIResponse 将 AlibabaAliqinFcVoiceRecordGeturlAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcVoiceRecordGeturlAPIResponse(v *AlibabaAliqinFcVoiceRecordGeturlAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcVoiceRecordGeturlAPIResponse.Put(v)
}
