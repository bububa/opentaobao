package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemUpdateListingTmall taobao.item.update.listing天猫分流
// taobao.item.update.listing.tmall
//
// * 单个商品上架&lt;br/&gt;* 输入的num_iid必须属于当前会话用户
func TaobaoItemUpdateListingTmall(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemUpdateListingTmallAPIRequest, resp *tbitem.TaobaoItemUpdateListingTmallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
