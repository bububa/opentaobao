package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频草稿信息同步 APIRequest
alibaba.alihouse.newhome.video.sync

接收视频信息记录
*/
type AlibabaAlihouseNewhomeVideoSyncRequest struct {
    model.Params

    // 草稿对下
    video   *VideoDraftDto 

}

func NewAlibabaAlihouseNewhomeVideoSyncRequest() *AlibabaAlihouseNewhomeVideoSyncRequest{
    return &AlibabaAlihouseNewhomeVideoSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeVideoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.video.sync"
}

func (r AlibabaAlihouseNewhomeVideoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeVideoSyncRequest) SetVideo(video *VideoDraftDto) error {
    r.video = video
    r.Set("video", video)
    return nil
}

func (r AlibabaAlihouseNewhomeVideoSyncRequest) GetVideo() *VideoDraftDto {
    return r.video
}

