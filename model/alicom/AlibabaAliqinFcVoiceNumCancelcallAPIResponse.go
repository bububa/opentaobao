package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcVoiceNumCancelcallAPIResponse 取消呼叫 API返回值
// alibaba.aliqin.fc.voice.num.cancelcall
//
// 当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
type AlibabaAliqinFcVoiceNumCancelcallAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcVoiceNumCancelcallAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcVoiceNumCancelcallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcVoiceNumCancelcallAPIResponseModel).Reset()
}

// AlibabaAliqinFcVoiceNumCancelcallAPIResponseModel is 取消呼叫 成功返回结果
type AlibabaAliqinFcVoiceNumCancelcallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_num_cancelcall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcVoiceNumCancelcallBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcVoiceNumCancelcallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcVoiceNumCancelcallAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcVoiceNumCancelcallAPIResponse)
	},
}

// GetAlibabaAliqinFcVoiceNumCancelcallAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcVoiceNumCancelcallAPIResponse
func GetAlibabaAliqinFcVoiceNumCancelcallAPIResponse() *AlibabaAliqinFcVoiceNumCancelcallAPIResponse {
	return poolAlibabaAliqinFcVoiceNumCancelcallAPIResponse.Get().(*AlibabaAliqinFcVoiceNumCancelcallAPIResponse)
}

// ReleaseAlibabaAliqinFcVoiceNumCancelcallAPIResponse 将 AlibabaAliqinFcVoiceNumCancelcallAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcVoiceNumCancelcallAPIResponse(v *AlibabaAliqinFcVoiceNumCancelcallAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcVoiceNumCancelcallAPIResponse.Put(v)
}
