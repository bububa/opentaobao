package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotSmsSendAPIResponse 物联网短信发送 API返回值
// alibaba.aliqin.fc.iot.sms.send
//
// 发送物联网短信，只允许使用物联网短信模板
type AlibabaAliqinFcIotSmsSendAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotSmsSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotSmsSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotSmsSendAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotSmsSendAPIResponseModel is 物联网短信发送 成功返回结果
type AlibabaAliqinFcIotSmsSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_sms_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAliqinFcIotSmsSendBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotSmsSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotSmsSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotSmsSendAPIResponse)
	},
}

// GetAlibabaAliqinFcIotSmsSendAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotSmsSendAPIResponse
func GetAlibabaAliqinFcIotSmsSendAPIResponse() *AlibabaAliqinFcIotSmsSendAPIResponse {
	return poolAlibabaAliqinFcIotSmsSendAPIResponse.Get().(*AlibabaAliqinFcIotSmsSendAPIResponse)
}

// ReleaseAlibabaAliqinFcIotSmsSendAPIResponse 将 AlibabaAliqinFcIotSmsSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotSmsSendAPIResponse(v *AlibabaAliqinFcIotSmsSendAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotSmsSendAPIResponse.Put(v)
}
