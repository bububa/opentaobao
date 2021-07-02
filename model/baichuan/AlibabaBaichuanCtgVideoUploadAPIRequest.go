package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgVideoUploadAPIRequest 提供优酷的短视频入淘API API请求
// alibaba.baichuan.ctg.video.upload
//
// 提供优酷的短视频入淘API
type AlibabaBaichuanCtgVideoUploadAPIRequest struct {
	model.Params
	// app
	_app string
	// type
	_type string
	// 优酷道长绑定的淘宝账号ID
	_tbUid string
	// 视频VID，若为多个视频，则支持分组上传多个VID
	_videoId string
	// 作者名称
	_ownerName string
	// 发布时间
	_publishTime string
	// 上传时间
	_uploadTime string
	// 视频标题
	_videoTitle string
	// 视频描述
	_videoInfo string
	// 视频的分类ID，目前是优酷的分类ID
	_videoCategory string
	// 视频标签
	_videoTag string
	// 视频的平台来源，如，优酷
	_source string
}

// NewAlibabaBaichuanCtgVideoUploadRequest 初始化AlibabaBaichuanCtgVideoUploadAPIRequest对象
func NewAlibabaBaichuanCtgVideoUploadRequest() *AlibabaBaichuanCtgVideoUploadAPIRequest {
	return &AlibabaBaichuanCtgVideoUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.video.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApp is App Setter
// app
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetApp(_app string) error {
	r._app = _app
	r.Set("app", _app)
	return nil
}

// GetApp App Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetApp() string {
	return r._app
}

// SetType is Type Setter
// type
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetType() string {
	return r._type
}

// SetTbUid is TbUid Setter
// 优酷道长绑定的淘宝账号ID
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetTbUid(_tbUid string) error {
	r._tbUid = _tbUid
	r.Set("tb_uid", _tbUid)
	return nil
}

// GetTbUid TbUid Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetTbUid() string {
	return r._tbUid
}

// SetVideoId is VideoId Setter
// 视频VID，若为多个视频，则支持分组上传多个VID
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetVideoId(_videoId string) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetVideoId() string {
	return r._videoId
}

// SetOwnerName is OwnerName Setter
// 作者名称
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetOwnerName(_ownerName string) error {
	r._ownerName = _ownerName
	r.Set("owner_name", _ownerName)
	return nil
}

// GetOwnerName OwnerName Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetOwnerName() string {
	return r._ownerName
}

// SetPublishTime is PublishTime Setter
// 发布时间
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetPublishTime(_publishTime string) error {
	r._publishTime = _publishTime
	r.Set("publish_time", _publishTime)
	return nil
}

// GetPublishTime PublishTime Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetPublishTime() string {
	return r._publishTime
}

// SetUploadTime is UploadTime Setter
// 上传时间
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetUploadTime(_uploadTime string) error {
	r._uploadTime = _uploadTime
	r.Set("upload_time", _uploadTime)
	return nil
}

// GetUploadTime UploadTime Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetUploadTime() string {
	return r._uploadTime
}

// SetVideoTitle is VideoTitle Setter
// 视频标题
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetVideoTitle(_videoTitle string) error {
	r._videoTitle = _videoTitle
	r.Set("video_title", _videoTitle)
	return nil
}

// GetVideoTitle VideoTitle Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetVideoTitle() string {
	return r._videoTitle
}

// SetVideoInfo is VideoInfo Setter
// 视频描述
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetVideoInfo(_videoInfo string) error {
	r._videoInfo = _videoInfo
	r.Set("video_info", _videoInfo)
	return nil
}

// GetVideoInfo VideoInfo Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetVideoInfo() string {
	return r._videoInfo
}

// SetVideoCategory is VideoCategory Setter
// 视频的分类ID，目前是优酷的分类ID
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetVideoCategory(_videoCategory string) error {
	r._videoCategory = _videoCategory
	r.Set("video_category", _videoCategory)
	return nil
}

// GetVideoCategory VideoCategory Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetVideoCategory() string {
	return r._videoCategory
}

// SetVideoTag is VideoTag Setter
// 视频标签
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetVideoTag(_videoTag string) error {
	r._videoTag = _videoTag
	r.Set("video_tag", _videoTag)
	return nil
}

// GetVideoTag VideoTag Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetVideoTag() string {
	return r._videoTag
}

// SetSource is Source Setter
// 视频的平台来源，如，优酷
func (r *AlibabaBaichuanCtgVideoUploadAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaBaichuanCtgVideoUploadAPIRequest) GetSource() string {
	return r._source
}
