package ott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取播放串地址 APIRequest
youku.ott.playservice.getplayurl

获取播放串地址服务
*/
type YoukuOttPlayserviceGetplayurlRequest struct {
    model.Params

    // 优酷账号登录态
    yktk   string 

    // 账号登录态
    havanaToken   string 

    // 系统信息
    systemInfo   string 

    // 视频ID
    videoId   int64 

}

func NewYoukuOttPlayserviceGetplayurlRequest() *YoukuOttPlayserviceGetplayurlRequest{
    return &YoukuOttPlayserviceGetplayurlRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttPlayserviceGetplayurlRequest) GetApiMethodName() string {
    return "youku.ott.playservice.getplayurl"
}

func (r YoukuOttPlayserviceGetplayurlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttPlayserviceGetplayurlRequest) SetYktk(yktk string) error {
    r.yktk = yktk
    r.Set("yktk", yktk)
    return nil
}

func (r YoukuOttPlayserviceGetplayurlRequest) GetYktk() string {
    return r.yktk
}

func (r *YoukuOttPlayserviceGetplayurlRequest) SetHavanaToken(havanaToken string) error {
    r.havanaToken = havanaToken
    r.Set("havana_token", havanaToken)
    return nil
}

func (r YoukuOttPlayserviceGetplayurlRequest) GetHavanaToken() string {
    return r.havanaToken
}

func (r *YoukuOttPlayserviceGetplayurlRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

func (r YoukuOttPlayserviceGetplayurlRequest) GetSystemInfo() string {
    return r.systemInfo
}

func (r *YoukuOttPlayserviceGetplayurlRequest) SetVideoId(videoId int64) error {
    r.videoId = videoId
    r.Set("video_id", videoId)
    return nil
}

func (r YoukuOttPlayserviceGetplayurlRequest) GetVideoId() int64 {
    return r.videoId
}

