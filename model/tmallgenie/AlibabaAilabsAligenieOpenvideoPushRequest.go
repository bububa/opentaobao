package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容库视频分集数据推送接口 APIRequest
alibaba.ailabs.aligenie.openvideo.push

天猫精灵内容库视频分集数据推送接口
*/
type AlibabaAilabsAligenieOpenvideoPushRequest struct {
    model.Params

    // 待推送的视频数据
    videos   []RawSingleVideo 

}

func NewAlibabaAilabsAligenieOpenvideoPushRequest() *AlibabaAilabsAligenieOpenvideoPushRequest{
    return &AlibabaAilabsAligenieOpenvideoPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieOpenvideoPushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.openvideo.push"
}

func (r AlibabaAilabsAligenieOpenvideoPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieOpenvideoPushRequest) SetVideos(videos []RawSingleVideo) error {
    r.videos = videos
    r.Set("videos", videos)
    return nil
}

func (r AlibabaAilabsAligenieOpenvideoPushRequest) GetVideos() []RawSingleVideo {
    return r.videos
}

