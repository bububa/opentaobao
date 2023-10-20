package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcVoiceNumDoublecallAPIResponse 多方通话 API返回值
// alibaba.aliqin.fc.voice.num.doublecall
//
// 根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
type AlibabaAliqinFcVoiceNumDoublecallAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcVoiceNumDoublecallAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcVoiceNumDoublecallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcVoiceNumDoublecallAPIResponseModel).Reset()
}

// AlibabaAliqinFcVoiceNumDoublecallAPIResponseModel is 多方通话 成功返回结果
type AlibabaAliqinFcVoiceNumDoublecallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_num_doublecall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回参数
	Result *AlibabaAliqinFcVoiceNumDoublecallBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcVoiceNumDoublecallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcVoiceNumDoublecallAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcVoiceNumDoublecallAPIResponse)
	},
}

// GetAlibabaAliqinFcVoiceNumDoublecallAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcVoiceNumDoublecallAPIResponse
func GetAlibabaAliqinFcVoiceNumDoublecallAPIResponse() *AlibabaAliqinFcVoiceNumDoublecallAPIResponse {
	return poolAlibabaAliqinFcVoiceNumDoublecallAPIResponse.Get().(*AlibabaAliqinFcVoiceNumDoublecallAPIResponse)
}

// ReleaseAlibabaAliqinFcVoiceNumDoublecallAPIResponse 将 AlibabaAliqinFcVoiceNumDoublecallAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcVoiceNumDoublecallAPIResponse(v *AlibabaAliqinFcVoiceNumDoublecallAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcVoiceNumDoublecallAPIResponse.Put(v)
}
