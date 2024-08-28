package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupMobilediscountDelete 批量删除adgroup的移动溢价
// taobao.simba.adgroup.mobilediscount.delete
//
// 批量删除adgroup的移动溢价
func TaobaoSimbaAdgroupMobilediscountDelete(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest, resp *simba.TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
