package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音频相关鉴权接口 APIRequest
alibaba.interact.media.audio

新音频包的鉴权接口
*/
type AlibabaInteractMediaAudioRequest struct {
    model.Params

    // 系统自动生成
    id   string 

}

func NewAlibabaInteractMediaAudioRequest() *AlibabaInteractMediaAudioRequest{
    return &AlibabaInteractMediaAudioRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractMediaAudioRequest) GetApiMethodName() string {
    return "alibaba.interact.media.audio"
}

func (r AlibabaInteractMediaAudioRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractMediaAudioRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaInteractMediaAudioRequest) GetId() string {
    return r.id
}

