package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemSimpleschemaUpdate 天猫简化编辑商品
// tmall.item.simpleschema.update
//
// 国外大商家天猫简化编辑商品
func TmallItemSimpleschemaUpdate(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemSimpleschemaUpdateAPIRequest, resp *tbitem.TmallItemSimpleschemaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
