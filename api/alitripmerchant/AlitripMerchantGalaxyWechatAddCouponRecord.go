package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechataddcouponrecord 星河-记录用户微信优惠券领取记录
// alitrip.merchant.galaxy.wechat.add.coupon.record
//
// 雅高小程序添加优惠券实例
func Alitripmerchantgalaxywechataddcouponrecord(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechataddcouponrecordAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechataddcouponrecordAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechataddcouponrecordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
