package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoBanamadpcItemEditRender 编辑商品发布页
// taobao.banamadpc.item.edit.render
//
// 巴拿马供应商通过此接口获取编辑商品发布页
func TaobaoBanamadpcItemEditRender(ctx context.Context, clt *core.SDKClient, req *product.TaobaoBanamadpcItemEditRenderAPIRequest, resp *product.TaobaoBanamadpcItemEditRenderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
