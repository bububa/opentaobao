package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcVoiceNumSinglecallAPIResponse 语音通知 API返回值
// alibaba.aliqin.fc.voice.num.singlecall
//
// 向指定手机号码发起单向呼叫，播放指定的语音文件内容。使用前需要在阿里大于管理中心添加去电显示号码与语音文件。
type AlibabaAliqinFcVoiceNumSinglecallAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcVoiceNumSinglecallAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcVoiceNumSinglecallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcVoiceNumSinglecallAPIResponseModel).Reset()
}

// AlibabaAliqinFcVoiceNumSinglecallAPIResponseModel is 语音通知 成功返回结果
type AlibabaAliqinFcVoiceNumSinglecallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_num_singlecall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAliqinFcVoiceNumSinglecallBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcVoiceNumSinglecallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcVoiceNumSinglecallAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcVoiceNumSinglecallAPIResponse)
	},
}

// GetAlibabaAliqinFcVoiceNumSinglecallAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcVoiceNumSinglecallAPIResponse
func GetAlibabaAliqinFcVoiceNumSinglecallAPIResponse() *AlibabaAliqinFcVoiceNumSinglecallAPIResponse {
	return poolAlibabaAliqinFcVoiceNumSinglecallAPIResponse.Get().(*AlibabaAliqinFcVoiceNumSinglecallAPIResponse)
}

// ReleaseAlibabaAliqinFcVoiceNumSinglecallAPIResponse 将 AlibabaAliqinFcVoiceNumSinglecallAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcVoiceNumSinglecallAPIResponse(v *AlibabaAliqinFcVoiceNumSinglecallAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcVoiceNumSinglecallAPIResponse.Put(v)
}
