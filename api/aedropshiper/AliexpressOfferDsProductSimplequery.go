package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressOfferDsProductSimplequery Dropshipper查询单个商品的简易信息
// aliexpress.offer.ds.product.simplequery
//
// 提供给Dropshipper的通过商品ID查找商品简易信息（包括SKU-价格/库存、产品状态等）的接口，只有特定买家可以使用
func AliexpressOfferDsProductSimplequery(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressOfferDsProductSimplequeryAPIRequest, resp *aedropshiper.AliexpressOfferDsProductSimplequeryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
