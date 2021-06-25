package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
sdk信息回调 APIRequest
taobao.top.sdk.feedback.upload

sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
*/
type TaobaoTopSdkFeedbackUploadRequest struct {
    model.Params

    // 1、回传加密信息
    type   string 

    // 具体内容，json形式
    content   string 

}

func NewTaobaoTopSdkFeedbackUploadRequest() *TaobaoTopSdkFeedbackUploadRequest{
    return &TaobaoTopSdkFeedbackUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTopSdkFeedbackUploadRequest) GetApiMethodName() string {
    return "taobao.top.sdk.feedback.upload"
}

func (r TaobaoTopSdkFeedbackUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTopSdkFeedbackUploadRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoTopSdkFeedbackUploadRequest) GetType() string {
    return r.type
}

func (r *TaobaoTopSdkFeedbackUploadRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoTopSdkFeedbackUploadRequest) GetContent() string {
    return r.content
}

