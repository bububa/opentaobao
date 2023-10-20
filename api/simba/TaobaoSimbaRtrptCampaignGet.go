package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRtrptCampaignGet 获取推广计划实时报表数据
// taobao.simba.rtrpt.campaign.get
//
// 获取推广计划实时报表数据
func TaobaoSimbaRtrptCampaignGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCampaignGetAPIRequest, resp *simba.TaobaoSimbaRtrptCampaignGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
