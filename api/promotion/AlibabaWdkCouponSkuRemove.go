package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponSkuRemove 优惠券商品删除
// alibaba.wdk.coupon.sku.remove
//
// 优惠券商品删除
func AlibabaWdkCouponSkuRemove(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaWdkCouponSkuRemoveAPIRequest, resp *promotion.AlibabaWdkCouponSkuRemoveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
