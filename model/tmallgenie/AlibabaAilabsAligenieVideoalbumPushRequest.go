package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容库视频合辑数据推送接口 API请求
alibaba.ailabs.aligenie.videoalbum.push

三方内容合作厂商可将视频辑数据通过此接口推送至天猫精灵内容库接入中，供天猫精灵使用
*/
type AlibabaAilabsAligenieVideoalbumPushRequest struct {
    model.Params
    // 视频合辑数据
    param1   []RawVideoAlbum
}

// 初始化AlibabaAilabsAligenieVideoalbumPushRequest对象
func NewAlibabaAilabsAligenieVideoalbumPushRequest() *AlibabaAilabsAligenieVideoalbumPushRequest{
    return &AlibabaAilabsAligenieVideoalbumPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieVideoalbumPushRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.videoalbum.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieVideoalbumPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param1 Setter
// 视频合辑数据
func (r *AlibabaAilabsAligenieVideoalbumPushRequest) SetParam1(param1 []RawVideoAlbum) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaAilabsAligenieVideoalbumPushRequest) GetParam1() []RawVideoAlbum {
    return r.param1
}
