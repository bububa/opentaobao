package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptAdgroupbaseGet 推广组基础报表数据对象
// taobao.simba.rpt.adgroupbase.get
//
// 推广组基础报表数据对象
func TaobaoSimbaRptAdgroupbaseGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupbaseGetAPIRequest, resp *simba.TaobaoSimbaRptAdgroupbaseGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
