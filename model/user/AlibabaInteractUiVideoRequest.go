package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 API请求
alibaba.interact.ui.video

Weex页面播放视频
*/
type AlibabaInteractUiVideoRequest struct {
    model.Params
    // 仅作鉴权使用，没有实际数据传输
    unnamed   string
}

// 初始化AlibabaInteractUiVideoRequest对象
func NewAlibabaInteractUiVideoRequest() *AlibabaInteractUiVideoRequest{
    return &AlibabaInteractUiVideoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractUiVideoRequest) GetApiMethodName() string {
    return "alibaba.interact.ui.video"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractUiVideoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Unnamed Setter
// 仅作鉴权使用，没有实际数据传输
func (r *AlibabaInteractUiVideoRequest) SetUnnamed(unnamed string) error {
    r.unnamed = unnamed
    r.Set("unnamed", unnamed)
    return nil
}

// Unnamed Getter
func (r AlibabaInteractUiVideoRequest) GetUnnamed() string {
    return r.unnamed
}
