package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRtrptAdgroupGet 获取推广组实时报表数据
// taobao.simba.rtrpt.adgroup.get
//
// 获取推广组实时报表数据
func TaobaoSimbaRtrptAdgroupGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRtrptAdgroupGetAPIRequest, resp *simba.TaobaoSimbaRtrptAdgroupGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
