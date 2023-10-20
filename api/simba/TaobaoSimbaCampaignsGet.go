package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignsGet 取得一组推广计划
// taobao.simba.campaigns.get
//
// 取得一个客户的推广计划；
func TaobaoSimbaCampaignsGet(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignsGetAPIRequest, resp *simba.TaobaoSimbaCampaignsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
