package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCampaignRealtime 查询某计划实时列表
// taobao.onebp.dkx.report.report.campaign.realtime
//
// 查询某计划实时列表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;log_date_list&#34;: [&#34;2021-09-09&#34;],     &#34;campaign_id_list&#34;: [2821811599]}
func TaobaoOnebpDkxReportReportCampaignRealtime(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
