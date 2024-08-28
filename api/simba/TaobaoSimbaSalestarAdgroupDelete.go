package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarAdgroupDelete (新)销量明星删除推广单元接口
// taobao.simba.salestar.adgroup.delete
//
// 删除一个推广组
func TaobaoSimbaSalestarAdgroupDelete(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupDeleteAPIRequest, resp *simba.TaobaoSimbaSalestarAdgroupDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
