package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptAdgroupkeywordeffectGet 推广组下的词效果报表数据查询(明细数据不分类型查询)
// taobao.simba.rpt.adgroupkeywordeffect.get
//
// 推广组下的词效果报表数据查询(明细数据不分类型查询)
func TaobaoSimbaRptAdgroupkeywordeffectGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest, resp *simba.TaobaoSimbaRptAdgroupkeywordeffectGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
