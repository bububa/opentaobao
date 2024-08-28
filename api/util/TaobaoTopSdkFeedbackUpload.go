package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopSdkFeedbackUpload sdk信息回调
// taobao.top.sdk.feedback.upload
//
// sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
func TaobaoTopSdkFeedbackUpload(ctx context.Context, clt *core.SDKClient, req *util.TaobaoTopSdkFeedbackUploadAPIRequest, resp *util.TaobaoTopSdkFeedbackUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
