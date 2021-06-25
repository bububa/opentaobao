package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
sdk信息回调 APIResponse
taobao.top.sdk.feedback.upload

sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
*/
type TaobaoTopSdkFeedbackUploadAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTopSdkFeedbackUploadResponse `json:"taobao_top_sdk_feedback_upload_response,omitempty"`
}

type TaobaoTopSdkFeedbackUploadResponse struct {

    // 控制回传间隔（单位：秒）
    UploadInterval   int64 `json:"upload_interval,omitempty"`

}
