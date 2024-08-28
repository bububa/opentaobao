package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptAdgroupkeywordbaseGet 推广组下的词基础报表数据查询(明细数据不分类型查询)
// taobao.simba.rpt.adgroupkeywordbase.get
//
// 推广组下的词基础报表数据查询(明细数据不分类型查询)
func TaobaoSimbaRptAdgroupkeywordbaseGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupkeywordbaseGetAPIRequest, resp *simba.TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
