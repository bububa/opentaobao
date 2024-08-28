package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupAdd 创建一个推广组
// taobao.simba.adgroup.add
//
// 创建一个推广组
func TaobaoSimbaAdgroupAdd(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupAddAPIRequest, resp *simba.TaobaoSimbaAdgroupAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
