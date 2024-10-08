package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrInventoryUpdate 门店业务同步库存
// tmall.nr.inventory.update
//
// 用于商家每日同步更新门店库存
func TmallNrInventoryUpdate(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrInventoryUpdateAPIRequest, resp *tmallnr.TmallNrInventoryUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
