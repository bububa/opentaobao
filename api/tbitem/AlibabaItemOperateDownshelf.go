package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemOperateDownshelf 商品下架
// alibaba.item.operate.downshelf
//
// 商品下架
func AlibabaItemOperateDownshelf(ctx context.Context, clt *core.SDKClient, req *tbitem.AlibabaItemOperateDownshelfAPIRequest, resp *tbitem.AlibabaItemOperateDownshelfAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
