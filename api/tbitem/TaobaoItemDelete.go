package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemDelete 删除单条商品
// taobao.item.delete
//
// 删除单条商品
func TaobaoItemDelete(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemDeleteAPIRequest, resp *tbitem.TaobaoItemDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
