package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse 视频专辑场景接入接口 API返回值
// alibaba.ailabs.aligenie.openvideoalbum.scenepush
//
// 视频专辑场景接入接口
type AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponseModel is 视频专辑场景接入接口 成功返回结果
type AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideoalbum_scenepush_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 状态码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetMsg = ""
	m.RetCode = 0
}

var poolAlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse)
	},
}

// GetAlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse
func GetAlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse() *AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse {
	return poolAlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse.Get().(*AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse)
}

// ReleaseAlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse 将 AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse(v *AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse.Put(v)
}
