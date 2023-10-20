package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCampaignbaseGet 推广计划报表基础数据对象
// taobao.simba.rpt.campaignbase.get
//
// 推广计划报表基础数据对象
func TaobaoSimbaRptCampaignbaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptCampaignbaseGetAPIRequest, resp *simba.TaobaoSimbaRptCampaignbaseGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
