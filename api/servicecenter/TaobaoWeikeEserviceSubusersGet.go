package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoWeikeEserviceSubusersGet 客服外包订单分配的商家子账号列表
// taobao.weike.eservice.subusers.get
//
// 获取客服外包订单分配的商家子账号列表，以及授权状态
func TaobaoWeikeEserviceSubusersGet(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceSubusersGetAPIRequest, resp *servicecenter.TaobaoWeikeEserviceSubusersGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
