package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoShopSellerGet 卖家店铺基础信息查询
// taobao.shop.seller.get
//
// 获取卖家店铺的基本信息
func TaobaoShopSellerGet(ctx context.Context, clt *core.SDKClient, req *shop.TaobaoShopSellerGetAPIRequest, resp *shop.TaobaoShopSellerGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
