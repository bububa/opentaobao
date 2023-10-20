package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingcouponqueryitems 查询优惠券活动下的商品
// alibaba.hm.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
func Alibabahmmarketingcouponqueryitems(clt *core.SDKClient, req *wdk.AlibabahmmarketingcouponqueryitemsAPIRequest, session string) (*wdk.AlibabahmmarketingcouponqueryitemsAPIResponse, error) {
	var resp wdk.AlibabahmmarketingcouponqueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
