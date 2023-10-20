package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingcouponadditem 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
// alibaba.wdk.marketing.coupon.additem
//
// 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
// 如果是商品券，则添加的商品为券适用的商品；
// 如果是品类券，则添加的商品为券排除的商品；
func Alibabawdkmarketingcouponadditem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingcouponadditemAPIRequest, session string) (*wdk.AlibabawdkmarketingcouponadditemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingcouponadditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
