package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCampaignRealtime 查询某计划实时列表
// taobao.onebp.dkx.report.report.campaign.realtime
//
// 查询某计划实时列表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"log_date_list": ["2021-09-09"],     "campaign_id_list": [2821811599]}
func TaobaoOnebpDkxReportReportCampaignRealtime(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
