package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemImgDelete 删除商品图片
// taobao.item.img.delete
//
// 删除商品图片
func TaobaoItemImgDelete(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemImgDeleteAPIRequest, resp *tbitem.TaobaoItemImgDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
