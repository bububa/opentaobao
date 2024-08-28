package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupsItemExist 商品是否推广
// taobao.simba.adgroups.item.exist
//
// 判断在一个推广计划中是否已经推广了一个商品
func TaobaoSimbaAdgroupsItemExist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupsItemExistAPIRequest, resp *simba.TaobaoSimbaAdgroupsItemExistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
