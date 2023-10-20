package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpenvideoScenepushAPIResponse 视频单集场景接入API API返回值
// alibaba.ailabs.aligenie.openvideo.scenepush
//
// 视频单集场景接入API
type AlibabaAilabsAligenieOpenvideoScenepushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieOpenvideoScenepushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpenvideoScenepushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieOpenvideoScenepushAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieOpenvideoScenepushAPIResponseModel is 视频单集场景接入API 成功返回结果
type AlibabaAilabsAligenieOpenvideoScenepushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideo_scenepush_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 状态码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpenvideoScenepushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetMsg = ""
	m.RetCode = 0
}

var poolAlibabaAilabsAligenieOpenvideoScenepushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieOpenvideoScenepushAPIResponse)
	},
}

// GetAlibabaAilabsAligenieOpenvideoScenepushAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieOpenvideoScenepushAPIResponse
func GetAlibabaAilabsAligenieOpenvideoScenepushAPIResponse() *AlibabaAilabsAligenieOpenvideoScenepushAPIResponse {
	return poolAlibabaAilabsAligenieOpenvideoScenepushAPIResponse.Get().(*AlibabaAilabsAligenieOpenvideoScenepushAPIResponse)
}

// ReleaseAlibabaAilabsAligenieOpenvideoScenepushAPIResponse 将 AlibabaAilabsAligenieOpenvideoScenepushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieOpenvideoScenepushAPIResponse(v *AlibabaAilabsAligenieOpenvideoScenepushAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieOpenvideoScenepushAPIResponse.Put(v)
}
