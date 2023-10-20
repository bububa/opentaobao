package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCampaignCampaignReportpage 获取场景计划的报表数据
// taobao.onebp.dkx.campaign.campaign.reportpage
//
// 获取场景计划的报表数据。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxCampaignCampaignReportpage(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest, resp *scs.TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
