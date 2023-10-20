package damaiticklet

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damaiticklet"
)

// AlibabaDamaiMxOpengatewayScript 第三方剧本数据推送
// alibaba.damai.mx.opengateway.script
//
// 第三方剧本数据推送
func AlibabaDamaiMxOpengatewayScript(clt *core.SDKClient, req *damaiticklet.AlibabaDamaiMxOpengatewayScriptAPIRequest, resp *damaiticklet.AlibabaDamaiMxOpengatewayScriptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
