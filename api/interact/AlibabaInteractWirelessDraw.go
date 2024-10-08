package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractWirelessDraw 双11到店互动无线端抽奖接口鉴权
// alibaba.interact.wireless.draw
//
// 双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
func AlibabaInteractWirelessDraw(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractWirelessDrawAPIRequest, resp *interact.AlibabaInteractWirelessDrawAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
