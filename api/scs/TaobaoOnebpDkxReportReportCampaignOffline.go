package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCampaignOffline 查询某计划离线列表
// taobao.onebp.dkx.report.report.campaign.offline
//
// 查询某计划离线列表；
// 拓展流量查询：
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"launch_product_id_list":[101004013],"start_time":"2021-04-26","campaign_id_list":[134821085],"end_time":"2021-04-28","effect":15,}
// 非拓展流量查询：
// 入参2示例：{"start_time":"2021-09-08","campaign_id_list":[2821811599],"end_time":"2021-09-08","effect":15}
func TaobaoOnebpDkxReportReportCampaignOffline(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
