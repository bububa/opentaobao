package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryArea 地域报表查询
// taobao.universalbp.report.query.area
//
// 地域报表查询
func TaobaoUniversalbpReportQueryArea(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryAreaAPIRequest, resp *simba.TaobaoUniversalbpReportQueryAreaAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
