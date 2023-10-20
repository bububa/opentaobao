package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxreportreportcampaignoffline 查询某计划离线列表
// taobao.onebp.dkx.report.report.campaign.offline
//
// 查询某计划离线列表；
// 拓展流量查询：
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;launch_product_id_list&#34;:[101004013],&#34;start_time&#34;:&#34;2021-04-26&#34;,&#34;campaign_id_list&#34;:[134821085],&#34;end_time&#34;:&#34;2021-04-28&#34;,&#34;effect&#34;:15,}
// 非拓展流量查询：
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;campaign_id_list&#34;:[2821811599],&#34;end_time&#34;:&#34;2021-09-08&#34;,&#34;effect&#34;:15}
func Taobaoonebpdkxreportreportcampaignoffline(clt *core.SDKClient, req *scs.TaobaoonebpdkxreportreportcampaignofflineAPIRequest, session string) (*scs.TaobaoonebpdkxreportreportcampaignofflineAPIResponse, error) {
	var resp scs.TaobaoonebpdkxreportreportcampaignofflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
