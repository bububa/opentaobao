package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
sdk信息回调 API请求
taobao.top.sdk.feedback.upload

sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
*/
type TaobaoTopSdkFeedbackUploadRequest struct {
    model.Params
    // 1、回传加密信息
    _type   string
    // 具体内容，json形式
    _content   string
}

// 初始化TaobaoTopSdkFeedbackUploadRequest对象
func NewTaobaoTopSdkFeedbackUploadRequest() *TaobaoTopSdkFeedbackUploadRequest{
    return &TaobaoTopSdkFeedbackUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopSdkFeedbackUploadRequest) GetApiMethodName() string {
    return "taobao.top.sdk.feedback.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopSdkFeedbackUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 1、回传加密信息
func (r *TaobaoTopSdkFeedbackUploadRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoTopSdkFeedbackUploadRequest) GetType() string {
    return r._type
}
// Content Setter
// 具体内容，json形式
func (r *TaobaoTopSdkFeedbackUploadRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoTopSdkFeedbackUploadRequest) GetContent() string {
    return r._content
}
