package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupidsChangedGet 获取修改的推广组ID
// taobao.simba.adgroupids.changed.get
//
// 获取修改的推广组ID
func TaobaoSimbaAdgroupidsChangedGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupidsChangedGetAPIRequest, resp *simba.TaobaoSimbaAdgroupidsChangedGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
