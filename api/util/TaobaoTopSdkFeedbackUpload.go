package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
sdk信息回调 
taobao.top.sdk.feedback.upload

sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
*/
func TaobaoTopSdkFeedbackUpload(clt *core.SDKClient, req *util.TaobaoTopSdkFeedbackUploadRequest, session string) (*util.TaobaoTopSdkFeedbackUploadResponse, error) {
    var resp util.TaobaoTopSdkFeedbackUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
