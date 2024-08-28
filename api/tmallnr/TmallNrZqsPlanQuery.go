package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrZqsPlanQuery 周期送配送明细查询
// tmall.nr.zqs.plan.query
//
// 周期送配送明细查询
func TmallNrZqsPlanQuery(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrZqsPlanQueryAPIRequest, resp *tmallnr.TmallNrZqsPlanQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
