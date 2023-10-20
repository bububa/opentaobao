package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardantGatewayCallback 人卡关系回调
// alibaba.campus.guardant.gateway.callback
//
// 门禁供应商回调平台通知同步结果
func AlibabaCampusGuardantGatewayCallback(clt *core.SDKClient, req *campus.AlibabaCampusGuardantGatewayCallbackAPIRequest, resp *campus.AlibabaCampusGuardantGatewayCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
