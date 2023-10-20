package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoShopSellerGet 卖家店铺基础信息查询
// taobao.shop.seller.get
//
// 获取卖家店铺的基本信息
func TaobaoShopSellerGet(clt *core.SDKClient, req *shop.TaobaoShopSellerGetAPIRequest, resp *shop.TaobaoShopSellerGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
