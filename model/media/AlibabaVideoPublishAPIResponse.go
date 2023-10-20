package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaVideoPublishAPIResponse 发布视频 API返回值
// alibaba.video.publish
//
// 发布视频。
// 说明：发布视频5s后再查询视频信息，否则可能无法获取播放链接
type AlibabaVideoPublishAPIResponse struct {
	model.CommonResponse
	AlibabaVideoPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaVideoPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaVideoPublishAPIResponseModel).Reset()
}

// AlibabaVideoPublishAPIResponseModel is 发布视频 成功返回结果
type AlibabaVideoPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_video_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 视频id
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaVideoPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.VideoId = 0
}

var poolAlibabaVideoPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaVideoPublishAPIResponse)
	},
}

// GetAlibabaVideoPublishAPIResponse 从 sync.Pool 获取 AlibabaVideoPublishAPIResponse
func GetAlibabaVideoPublishAPIResponse() *AlibabaVideoPublishAPIResponse {
	return poolAlibabaVideoPublishAPIResponse.Get().(*AlibabaVideoPublishAPIResponse)
}

// ReleaseAlibabaVideoPublishAPIResponse 将 AlibabaVideoPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaVideoPublishAPIResponse(v *AlibabaVideoPublishAPIResponse) {
	v.Reset()
	poolAlibabaVideoPublishAPIResponse.Put(v)
}
