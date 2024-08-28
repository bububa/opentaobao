package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignsGet 取得一组推广计划
// taobao.simba.campaigns.get
//
// 取得一个客户的推广计划；
func TaobaoSimbaCampaignsGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignsGetAPIRequest, resp *simba.TaobaoSimbaCampaignsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
