package aiar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

// AlibabaAiArTmjlAppDetect 天猫精灵扫一扫入口的服务
// alibaba.ai.ar.tmjl.app.detect
//
// 天猫精灵扫一扫入口的图像检测服务
func AlibabaAiArTmjlAppDetect(ctx context.Context, clt *core.SDKClient, req *aiar.AlibabaAiArTmjlAppDetectAPIRequest, resp *aiar.AlibabaAiArTmjlAppDetectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
