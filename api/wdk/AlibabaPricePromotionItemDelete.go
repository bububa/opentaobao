package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaPricePromotionItemDelete 批量删除档期
// alibaba.price.promotion.item.delete
//
// 盒马帮批量删除档期商品
func AlibabaPricePromotionItemDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaPricePromotionItemDeleteAPIRequest, resp *wdk.AlibabaPricePromotionItemDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
