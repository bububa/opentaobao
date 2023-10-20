package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieVideoalbumPushAPIResponse 天猫精灵内容库视频合辑数据推送接口 API返回值
// alibaba.ailabs.aligenie.videoalbum.push
//
// 三方内容合作厂商可将视频辑数据通过此接口推送至天猫精灵内容库接入中，供天猫精灵使用
type AlibabaAilabsAligenieVideoalbumPushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieVideoalbumPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieVideoalbumPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieVideoalbumPushAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieVideoalbumPushAPIResponseModel is 天猫精灵内容库视频合辑数据推送接口 成功返回结果
type AlibabaAilabsAligenieVideoalbumPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_videoalbum_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 状态码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieVideoalbumPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetMsg = ""
	m.RetCode = 0
}

var poolAlibabaAilabsAligenieVideoalbumPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieVideoalbumPushAPIResponse)
	},
}

// GetAlibabaAilabsAligenieVideoalbumPushAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieVideoalbumPushAPIResponse
func GetAlibabaAilabsAligenieVideoalbumPushAPIResponse() *AlibabaAilabsAligenieVideoalbumPushAPIResponse {
	return poolAlibabaAilabsAligenieVideoalbumPushAPIResponse.Get().(*AlibabaAilabsAligenieVideoalbumPushAPIResponse)
}

// ReleaseAlibabaAilabsAligenieVideoalbumPushAPIResponse 将 AlibabaAilabsAligenieVideoalbumPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieVideoalbumPushAPIResponse(v *AlibabaAilabsAligenieVideoalbumPushAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieVideoalbumPushAPIResponse.Put(v)
}
