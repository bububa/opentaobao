package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaHourReportAdgroupGet 推广单元小时级别实时报表查询
// taobao.simba.hour.report.adgroup.get
//
// 推广单元小时级别实时报表查询
func TaobaoSimbaHourReportAdgroupGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaHourReportAdgroupGetAPIRequest, resp *simba.TaobaoSimbaHourReportAdgroupGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
