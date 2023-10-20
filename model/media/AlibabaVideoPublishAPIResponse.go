package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabavideopublishAPIResponse 发布视频 API返回值
// alibaba.video.publish
//
// 发布视频。
// 说明：发布视频5s后再查询视频信息，否则可能无法获取播放链接
type AlibabavideopublishAPIResponse struct {
	model.CommonResponse
	AlibabavideopublishAPIResponseModel
}

// AlibabavideopublishAPIResponseModel is 发布视频 成功返回结果
type AlibabavideopublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_video_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 视频id
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
}
