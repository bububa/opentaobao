package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemSimpleschemaAdd 天猫简化发布商品
// tmall.item.simpleschema.add
//
// 天猫简化版schema发布商品。
func TmallItemSimpleschemaAdd(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemSimpleschemaAddAPIRequest, resp *tbitem.TmallItemSimpleschemaAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
