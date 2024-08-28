package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoBanamadpcItemUpdate 编辑商品
// taobao.banamadpc.item.update
//
// 巴拿马供应商通过此接口编辑商品
func TaobaoBanamadpcItemUpdate(ctx context.Context, clt *core.SDKClient, req *product.TaobaoBanamadpcItemUpdateAPIRequest, resp *product.TaobaoBanamadpcItemUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
