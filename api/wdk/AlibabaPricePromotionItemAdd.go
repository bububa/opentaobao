package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaPricePromotionItemAdd 新增档期商品
// alibaba.price.promotion.item.add
//
// 批量新增档期活动商品
func AlibabaPricePromotionItemAdd(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaPricePromotionItemAddAPIRequest, resp *wdk.AlibabaPricePromotionItemAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
