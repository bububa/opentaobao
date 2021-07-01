package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容库视频分集数据推送接口 API请求
alibaba.ailabs.aligenie.openvideo.push

天猫精灵内容库视频分集数据推送接口
*/
type AlibabaAilabsAligenieOpenvideoPushAPIRequest struct {
    model.Params
    // 待推送的视频数据
    _videos   []RawSingleVideo
}

// 初始化AlibabaAilabsAligenieOpenvideoPushAPIRequest对象
func NewAlibabaAilabsAligenieOpenvideoPushRequest() *AlibabaAilabsAligenieOpenvideoPushAPIRequest{
    return &AlibabaAilabsAligenieOpenvideoPushAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpenvideoPushAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.openvideo.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpenvideoPushAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Videos Setter
// 待推送的视频数据
func (r *AlibabaAilabsAligenieOpenvideoPushAPIRequest) SetVideos(_videos []RawSingleVideo) error {
    r._videos = _videos
    r.Set("videos", _videos)
    return nil
}

// Videos Getter
func (r AlibabaAilabsAligenieOpenvideoPushAPIRequest) GetVideos() []RawSingleVideo {
    return r._videos
}
