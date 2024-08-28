package scs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportAccountOffline 获取账户历史报表
// taobao.onebp.dkx.report.report.account.offline
//
// 获取账户历史报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{     &#34;start_time&#34;: &#34;2021-07-24&#34;,     &#34;effect&#34;: 15,     &#34;end_time&#34;: &#34;2021-08-21&#34;,     &#34;strategy_scene&#34;:true,     &#34;unify_type&#34;:&#34;kuan&#34;,     &#34;bizCode&#34;:&#34;adStrategyDkx&#34;   }
func TaobaoOnebpDkxReportReportAccountOffline(ctx context.Context, clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportAccountOfflineAPIRequest, resp *scs.TaobaoOnebpDkxReportReportAccountOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
