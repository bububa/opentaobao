package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCampaignDaylist 获取计划分日报表
// taobao.onebp.dkx.report.report.campaign.daylist
//
// 获取计划分日报表，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe
func TaobaoOnebpDkxReportReportCampaignDaylist(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest, resp *scs.TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
