package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingCouponQueryitems 查询优惠券活动下的商品
// alibaba.wdk.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
func AlibabaWdkMarketingCouponQueryitems(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponQueryitemsAPIRequest, resp *wdk.AlibabaWdkMarketingCouponQueryitemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
