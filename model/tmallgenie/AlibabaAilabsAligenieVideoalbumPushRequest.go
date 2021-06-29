package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容库视频合辑数据推送接口 APIRequest
alibaba.ailabs.aligenie.videoalbum.push

三方内容合作厂商可将视频辑数据通过此接口推送至天猫精灵内容库接入中，供天猫精灵使用
*/
type AlibabaAilabsAligenieVideoalbumPushRequest struct {
    model.Params

    // 视频合辑数据
    param1   []RawVideoAlbum 

}

func NewAlibabaAilabsAligenieVideoalbumPushRequest() *AlibabaAilabsAligenieVideoalbumPushRequest{
    return &AlibabaAilabsAligenieVideoalbumPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieVideoalbumPushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.videoalbum.push"
}

func (r AlibabaAilabsAligenieVideoalbumPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieVideoalbumPushRequest) SetParam1(param1 []RawVideoAlbum) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaAilabsAligenieVideoalbumPushRequest) GetParam1() []RawVideoAlbum {
    return r.param1
}

