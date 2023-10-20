package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingcouponqueryitems 查询优惠券活动下的商品
// alibaba.wdk.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
func Alibabawdkmarketingcouponqueryitems(clt *core.SDKClient, req *wdk.AlibabawdkmarketingcouponqueryitemsAPIRequest, session string) (*wdk.AlibabawdkmarketingcouponqueryitemsAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingcouponqueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
