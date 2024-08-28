package damaiticklet

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damaiticklet"
)

// AlibabaDamaiMxOpengatewayScript 第三方剧本数据推送
// alibaba.damai.mx.opengateway.script
//
// 第三方剧本数据推送
func AlibabaDamaiMxOpengatewayScript(ctx context.Context, clt *core.SDKClient, req *damaiticklet.AlibabaDamaiMxOpengatewayScriptAPIRequest, resp *damaiticklet.AlibabaDamaiMxOpengatewayScriptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
