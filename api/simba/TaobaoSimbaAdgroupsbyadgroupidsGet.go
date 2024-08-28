package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupsbyadgroupidsGet 批量得到推广组
// taobao.simba.adgroupsbyadgroupids.get
//
// 批量得到推广组
func TaobaoSimbaAdgroupsbyadgroupidsGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest, resp *simba.TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
