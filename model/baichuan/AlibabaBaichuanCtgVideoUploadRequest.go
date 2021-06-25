package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提供优酷的短视频入淘API APIRequest
alibaba.baichuan.ctg.video.upload

提供优酷的短视频入淘API
*/
type AlibabaBaichuanCtgVideoUploadRequest struct {
    model.Params

    // app
    app   string 

    // type
    type   string 

    // 优酷道长绑定的淘宝账号ID
    tbUid   string 

    // 视频VID，若为多个视频，则支持分组上传多个VID
    videoId   string 

    // 作者名称
    ownerName   string 

    // 发布时间
    publishTime   string 

    // 上传时间
    uploadTime   string 

    // 视频标题
    videoTitle   string 

    // 视频描述
    videoInfo   string 

    // 视频的分类ID，目前是优酷的分类ID
    videoCategory   string 

    // 视频标签
    videoTag   string 

    // 视频的平台来源，如，优酷
    source   string 

}

func NewAlibabaBaichuanCtgVideoUploadRequest() *AlibabaBaichuanCtgVideoUploadRequest{
    return &AlibabaBaichuanCtgVideoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetApiMethodName() string {
    return "alibaba.baichuan.ctg.video.upload"
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaBaichuanCtgVideoUploadRequest) SetApp(app string) error {
    r.app = app
    r.Set("app", app)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetApp() string {
    return r.app
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetType() string {
    return r.type
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetTbUid(tbUid string) error {
    r.tbUid = tbUid
    r.Set("tb_uid", tbUid)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetTbUid() string {
    return r.tbUid
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoId(videoId string) error {
    r.videoId = videoId
    r.Set("video_id", videoId)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoId() string {
    return r.videoId
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetOwnerName(ownerName string) error {
    r.ownerName = ownerName
    r.Set("owner_name", ownerName)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetOwnerName() string {
    return r.ownerName
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetPublishTime(publishTime string) error {
    r.publishTime = publishTime
    r.Set("publish_time", publishTime)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetPublishTime() string {
    return r.publishTime
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetUploadTime(uploadTime string) error {
    r.uploadTime = uploadTime
    r.Set("upload_time", uploadTime)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetUploadTime() string {
    return r.uploadTime
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoTitle(videoTitle string) error {
    r.videoTitle = videoTitle
    r.Set("video_title", videoTitle)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoTitle() string {
    return r.videoTitle
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoInfo(videoInfo string) error {
    r.videoInfo = videoInfo
    r.Set("video_info", videoInfo)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoInfo() string {
    return r.videoInfo
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoCategory(videoCategory string) error {
    r.videoCategory = videoCategory
    r.Set("video_category", videoCategory)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoCategory() string {
    return r.videoCategory
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetVideoTag(videoTag string) error {
    r.videoTag = videoTag
    r.Set("video_tag", videoTag)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetVideoTag() string {
    return r.videoTag
}

func (r *AlibabaBaichuanCtgVideoUploadRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaBaichuanCtgVideoUploadRequest) GetSource() string {
    return r.source
}

