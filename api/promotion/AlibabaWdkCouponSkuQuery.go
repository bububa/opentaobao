package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponSkuQuery 优惠券商品查询
// alibaba.wdk.coupon.sku.query
//
// 优惠券商品查询
func AlibabaWdkCouponSkuQuery(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaWdkCouponSkuQueryAPIRequest, resp *promotion.AlibabaWdkCouponSkuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
