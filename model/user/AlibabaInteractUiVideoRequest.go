package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 APIRequest
alibaba.interact.ui.video

Weex页面播放视频
*/
type AlibabaInteractUiVideoRequest struct {
    model.Params

    // 仅作鉴权使用，没有实际数据传输
    unnamed   string 

}

func NewAlibabaInteractUiVideoRequest() *AlibabaInteractUiVideoRequest{
    return &AlibabaInteractUiVideoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractUiVideoRequest) GetApiMethodName() string {
    return "alibaba.interact.ui.video"
}

func (r AlibabaInteractUiVideoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractUiVideoRequest) SetUnnamed(unnamed string) error {
    r.unnamed = unnamed
    r.Set("unnamed", unnamed)
    return nil
}

func (r AlibabaInteractUiVideoRequest) GetUnnamed() string {
    return r.unnamed
}

