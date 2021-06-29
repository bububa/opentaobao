package ott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取播放串地址 API请求
youku.ott.playservice.getplayurl

获取播放串地址服务
*/
type YoukuOttPlayserviceGetplayurlRequest struct {
    model.Params
    // 优酷账号登录态
    _yktk   string
    // 账号登录态
    _havanaToken   string
    // 系统信息
    _systemInfo   string
    // 视频ID
    _videoId   int64
}

// 初始化YoukuOttPlayserviceGetplayurlRequest对象
func NewYoukuOttPlayserviceGetplayurlRequest() *YoukuOttPlayserviceGetplayurlRequest{
    return &YoukuOttPlayserviceGetplayurlRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttPlayserviceGetplayurlRequest) GetApiMethodName() string {
    return "youku.ott.playservice.getplayurl"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttPlayserviceGetplayurlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Yktk Setter
// 优酷账号登录态
func (r *YoukuOttPlayserviceGetplayurlRequest) SetYktk(_yktk string) error {
    r._yktk = _yktk
    r.Set("yktk", _yktk)
    return nil
}

// Yktk Getter
func (r YoukuOttPlayserviceGetplayurlRequest) GetYktk() string {
    return r._yktk
}
// HavanaToken Setter
// 账号登录态
func (r *YoukuOttPlayserviceGetplayurlRequest) SetHavanaToken(_havanaToken string) error {
    r._havanaToken = _havanaToken
    r.Set("havana_token", _havanaToken)
    return nil
}

// HavanaToken Getter
func (r YoukuOttPlayserviceGetplayurlRequest) GetHavanaToken() string {
    return r._havanaToken
}
// SystemInfo Setter
// 系统信息
func (r *YoukuOttPlayserviceGetplayurlRequest) SetSystemInfo(_systemInfo string) error {
    r._systemInfo = _systemInfo
    r.Set("system_info", _systemInfo)
    return nil
}

// SystemInfo Getter
func (r YoukuOttPlayserviceGetplayurlRequest) GetSystemInfo() string {
    return r._systemInfo
}
// VideoId Setter
// 视频ID
func (r *YoukuOttPlayserviceGetplayurlRequest) SetVideoId(_videoId int64) error {
    r._videoId = _videoId
    r.Set("video_id", _videoId)
    return nil
}

// VideoId Getter
func (r YoukuOttPlayserviceGetplayurlRequest) GetVideoId() int64 {
    return r._videoId
}
