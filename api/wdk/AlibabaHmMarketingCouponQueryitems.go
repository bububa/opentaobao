package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingCouponQueryitems 查询优惠券活动下的商品
// alibaba.hm.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
func AlibabaHmMarketingCouponQueryitems(clt *core.SDKClient, req *wdk.AlibabaHmMarketingCouponQueryitemsAPIRequest, session string) (*wdk.AlibabaHmMarketingCouponQueryitemsAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingCouponQueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
