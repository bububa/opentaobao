package lstspeacker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerStatusGetAPIResponse 音箱设备在线状态 API返回值
// alibaba.lst.speaker.status.get
//
// 音箱设备在线状态查询
type AlibabaLstSpeakerStatusGetAPIResponse struct {
	model.CommonResponse
	AlibabaLstSpeakerStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstSpeakerStatusGetAPIResponseModel).Reset()
}

// AlibabaLstSpeakerStatusGetAPIResponseModel is 音箱设备在线状态 成功返回结果
type AlibabaLstSpeakerStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaLstSpeakerStatusGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstSpeakerStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerStatusGetAPIResponse)
	},
}

// GetAlibabaLstSpeakerStatusGetAPIResponse 从 sync.Pool 获取 AlibabaLstSpeakerStatusGetAPIResponse
func GetAlibabaLstSpeakerStatusGetAPIResponse() *AlibabaLstSpeakerStatusGetAPIResponse {
	return poolAlibabaLstSpeakerStatusGetAPIResponse.Get().(*AlibabaLstSpeakerStatusGetAPIResponse)
}

// ReleaseAlibabaLstSpeakerStatusGetAPIResponse 将 AlibabaLstSpeakerStatusGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstSpeakerStatusGetAPIResponse(v *AlibabaLstSpeakerStatusGetAPIResponse) {
	v.Reset()
	poolAlibabaLstSpeakerStatusGetAPIResponse.Put(v)
}
