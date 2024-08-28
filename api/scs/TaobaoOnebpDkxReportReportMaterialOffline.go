package scs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportMaterialOffline 查询某计划分商品离线报表
// taobao.onebp.dkx.report.report.material.offline
//
// 查询某计划分商品离线报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-23&#34;,&#34;campaign_id_list&#34;:[2853805001],&#34;end_time&#34;:&#34;2021-09-24&#34;,&#34;effect&#34;: 15   }
func TaobaoOnebpDkxReportReportMaterialOffline(ctx context.Context, clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest, resp *scs.TaobaoOnebpDkxReportReportMaterialOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
