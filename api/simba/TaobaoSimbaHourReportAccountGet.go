package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaHourReportAccountGet 账户级别小时报表获取
// taobao.simba.hour.report.account.get
//
// 获取账户小时实时报表数据
func TaobaoSimbaHourReportAccountGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaHourReportAccountGetAPIRequest, resp *simba.TaobaoSimbaHourReportAccountGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
