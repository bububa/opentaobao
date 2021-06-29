package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提供优酷的短视频入淘API API请求
alibaba.baichuan.ctg.video.upload

提供优酷的短视频入淘API
*/
type AlibabaBaichuanCtgVideoUploadRequest struct {
    model.Params
    // app
    _app   string
    // type
    _type   string
    // 优酷道长绑定的淘宝账号ID
    _tbUid   string
    // 视频VID，若为多个视频，则支持分组上传多个VID
    _videoId   string
    // 作者名称
    _ownerName   string
    // 发布时间
    _publishTime   string
    // 上传时间
    _uploadTime   string
    // 视频标题
    _videoTitle   string
    // 视频描述
    _videoInfo   string
    // 视频的分类ID，目前是优酷的分类ID
    _videoCategory   string
    // 视频标签
    _videoTag   string
    // 视频的平台来源，如，优酷
    _source   string
}

// 初始化AlibabaBaichuanCtgVideoUploadRequest对象
func NewAlibabaBaichuanCtgVideoUploadRequest() *AlibabaBaichuanCtgVideoUploadRequest{
    return &AlibabaBaichuanCtgVideoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgVideoUploadRequest) GetApiMethodName() string {
    return "alibaba.baichuan.ctg.video.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgVideoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// App Setter
// app
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetApp(_app string) error {
    r._app = _app
    r.Set("app", _app)
    return nil
}

// App Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetApp() string {
    return r._app
}
// Type Setter
// type
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetType() string {
    return r._type
}
// TbUid Setter
// 优酷道长绑定的淘宝账号ID
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetTbUid(_tbUid string) error {
    r._tbUid = _tbUid
    r.Set("tb_uid", _tbUid)
    return nil
}

// TbUid Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetTbUid() string {
    return r._tbUid
}
// VideoId Setter
// 视频VID，若为多个视频，则支持分组上传多个VID
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoId(_videoId string) error {
    r._videoId = _videoId
    r.Set("video_id", _videoId)
    return nil
}

// VideoId Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoId() string {
    return r._videoId
}
// OwnerName Setter
// 作者名称
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetOwnerName(_ownerName string) error {
    r._ownerName = _ownerName
    r.Set("owner_name", _ownerName)
    return nil
}

// OwnerName Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetOwnerName() string {
    return r._ownerName
}
// PublishTime Setter
// 发布时间
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetPublishTime(_publishTime string) error {
    r._publishTime = _publishTime
    r.Set("publish_time", _publishTime)
    return nil
}

// PublishTime Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetPublishTime() string {
    return r._publishTime
}
// UploadTime Setter
// 上传时间
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetUploadTime(_uploadTime string) error {
    r._uploadTime = _uploadTime
    r.Set("upload_time", _uploadTime)
    return nil
}

// UploadTime Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetUploadTime() string {
    return r._uploadTime
}
// VideoTitle Setter
// 视频标题
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoTitle(_videoTitle string) error {
    r._videoTitle = _videoTitle
    r.Set("video_title", _videoTitle)
    return nil
}

// VideoTitle Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoTitle() string {
    return r._videoTitle
}
// VideoInfo Setter
// 视频描述
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoInfo(_videoInfo string) error {
    r._videoInfo = _videoInfo
    r.Set("video_info", _videoInfo)
    return nil
}

// VideoInfo Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoInfo() string {
    return r._videoInfo
}
// VideoCategory Setter
// 视频的分类ID，目前是优酷的分类ID
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoCategory(_videoCategory string) error {
    r._videoCategory = _videoCategory
    r.Set("video_category", _videoCategory)
    return nil
}

// VideoCategory Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoCategory() string {
    return r._videoCategory
}
// VideoTag Setter
// 视频标签
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoTag(_videoTag string) error {
    r._videoTag = _videoTag
    r.Set("video_tag", _videoTag)
    return nil
}

// VideoTag Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoTag() string {
    return r._videoTag
}
// Source Setter
// 视频的平台来源，如，优酷
func (r *AlibabaBaichuanCtgVideoUploadRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaBaichuanCtgVideoUploadRequest) GetSource() string {
    return r._source
}
