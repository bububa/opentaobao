package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaVideoQueryAPIResponse 查询视频信息 API返回值
// alibaba.video.query
//
// 查询视频信息
type AlibabaVideoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaVideoQueryAPIResponseModel
}

// AlibabaVideoQueryAPIResponseModel is 查询视频信息 成功返回结果
type AlibabaVideoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_video_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 视频标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 上传者id
	UploaderId string `json:"uploader_id,omitempty" xml:"uploader_id,omitempty"`
	// 视频封面
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 视频描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 视频比例
	AspectRatio string `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// 视频标签
	Tags string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 播放链接
	PlayUrl string `json:"play_url,omitempty" xml:"play_url,omitempty"`
	// 审核状态
	AuditState string `json:"audit_state,omitempty" xml:"audit_state,omitempty"`
	// 视频id
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
	// 上传时间
	UploadTime int64 `json:"upload_time,omitempty" xml:"upload_time,omitempty"`
	// 视频时长(s)
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 视频高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 视频宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}
