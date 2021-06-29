package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音频相关鉴权接口 API请求
alibaba.interact.media.audio

新音频包的鉴权接口
*/
type AlibabaInteractMediaAudioRequest struct {
    model.Params
    // 系统自动生成
    _id   string
}

// 初始化AlibabaInteractMediaAudioRequest对象
func NewAlibabaInteractMediaAudioRequest() *AlibabaInteractMediaAudioRequest{
    return &AlibabaInteractMediaAudioRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractMediaAudioRequest) GetApiMethodName() string {
    return "alibaba.interact.media.audio"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractMediaAudioRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 系统自动生成
func (r *AlibabaInteractMediaAudioRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaInteractMediaAudioRequest) GetId() string {
    return r._id
}
