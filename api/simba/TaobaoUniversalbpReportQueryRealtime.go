package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryRealtime 实时报表查询
// taobao.universalbp.report.query.realtime
//
// 实时报表查询
func TaobaoUniversalbpReportQueryRealtime(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryRealtimeAPIRequest, resp *simba.TaobaoUniversalbpReportQueryRealtimeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
