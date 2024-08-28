package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarAdgroupUpdate 销量明星更新一个推广组的信息
// taobao.simba.salestar.adgroup.update
//
// 更新一个推广组的信息，可以设置 是否上线
func TaobaoSimbaSalestarAdgroupUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupUpdateAPIRequest, resp *simba.TaobaoSimbaSalestarAdgroupUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
