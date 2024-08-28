package aiar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

// AlibabaAiArOpenPlatformDetect AR开发者平台marker图片检测服务
// alibaba.ai.ar.open.platform.detect
//
// AR开发者平台marker图片检测服务，给AR SDK 和 阿里火眼app使用
func AlibabaAiArOpenPlatformDetect(ctx context.Context, clt *core.SDKClient, req *aiar.AlibabaAiArOpenPlatformDetectAPIRequest, resp *aiar.AlibabaAiArOpenPlatformDetectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
