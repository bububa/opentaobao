package scs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportAccountRealtime 获取账户实时报表
// taobao.onebp.dkx.report.report.account.realtime
//
// 获取账户实时报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{     &#34;log_date_list&#34;: [  &#34;2021-09-23&#34;     ]   }
func TaobaoOnebpDkxReportReportAccountRealtime(ctx context.Context, clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest, resp *scs.TaobaoOnebpDkxReportReportAccountRealtimeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
