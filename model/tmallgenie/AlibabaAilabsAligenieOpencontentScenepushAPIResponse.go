package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpencontentScenepushAPIResponse 音频场景接入接口 API返回值
// alibaba.ailabs.aligenie.opencontent.scenepush
//
// 天猫精灵音频挂靠场景接入
type AlibabaAilabsAligenieOpencontentScenepushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieOpencontentScenepushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpencontentScenepushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieOpencontentScenepushAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieOpencontentScenepushAPIResponseModel is 音频场景接入接口 成功返回结果
type AlibabaAilabsAligenieOpencontentScenepushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_opencontent_scenepush_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 错误码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpencontentScenepushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetMsg = ""
	m.RetCode = 0
}

var poolAlibabaAilabsAligenieOpencontentScenepushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieOpencontentScenepushAPIResponse)
	},
}

// GetAlibabaAilabsAligenieOpencontentScenepushAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieOpencontentScenepushAPIResponse
func GetAlibabaAilabsAligenieOpencontentScenepushAPIResponse() *AlibabaAilabsAligenieOpencontentScenepushAPIResponse {
	return poolAlibabaAilabsAligenieOpencontentScenepushAPIResponse.Get().(*AlibabaAilabsAligenieOpencontentScenepushAPIResponse)
}

// ReleaseAlibabaAilabsAligenieOpencontentScenepushAPIResponse 将 AlibabaAilabsAligenieOpencontentScenepushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieOpencontentScenepushAPIResponse(v *AlibabaAilabsAligenieOpencontentScenepushAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieOpencontentScenepushAPIResponse.Put(v)
}
