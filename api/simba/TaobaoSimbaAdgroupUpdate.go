package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupUpdate 更新一个推广组的信息
// taobao.simba.adgroup.update
//
// 更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
func TaobaoSimbaAdgroupUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupUpdateAPIRequest, resp *simba.TaobaoSimbaAdgroupUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
