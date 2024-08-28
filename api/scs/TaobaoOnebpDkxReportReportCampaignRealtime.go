package scs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCampaignRealtime 查询某计划实时列表
// taobao.onebp.dkx.report.report.campaign.realtime
//
// 查询某计划实时列表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;log_date_list&#34;: [&#34;2021-09-09&#34;],     &#34;campaign_id_list&#34;: [2821811599]}
func TaobaoOnebpDkxReportReportCampaignRealtime(ctx context.Context, clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCampaignRealtimeAPIRequest, resp *scs.TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
