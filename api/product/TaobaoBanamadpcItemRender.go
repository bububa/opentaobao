package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoBanamadpcItemRender 新发商品发布页
// taobao.banamadpc.item.render
//
// 巴拿马供应商通过此接口新发商品发布页
func TaobaoBanamadpcItemRender(ctx context.Context, clt *core.SDKClient, req *product.TaobaoBanamadpcItemRenderAPIRequest, resp *product.TaobaoBanamadpcItemRenderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
