package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupMobilediscountUpdate 对推广组进行单独移动溢价
// taobao.simba.adgroup.mobilediscount.update
//
// 对推广组进行单独移动溢价
func TaobaoSimbaAdgroupMobilediscountUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest, resp *simba.TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
