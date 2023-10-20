package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponSkuAdd 优惠券商品增加
// alibaba.wdk.coupon.sku.add
//
// 优惠券商品增加
func AlibabaWdkCouponSkuAdd(clt *core.SDKClient, req *promotion.AlibabaWdkCouponSkuAddAPIRequest, resp *promotion.AlibabaWdkCouponSkuAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
