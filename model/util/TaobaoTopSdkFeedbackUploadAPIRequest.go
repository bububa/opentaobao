package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopSdkFeedbackUploadAPIRequest
sdk信息回调 API请求
taobao.top.sdk.feedback.upload

sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性 */
type TaobaoTopSdkFeedbackUploadAPIRequest struct {
	model.Params
	// 1、回传加密信息
	_type string
	// 具体内容，json形式
	_content string
}

// New
