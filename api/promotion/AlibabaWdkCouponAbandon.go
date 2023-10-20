package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcouponabandon 废券
// alibaba.wdk.coupon.abandon
//
// 优惠券废弃
func Alibabawdkcouponabandon(clt *core.SDKClient, req *promotion.AlibabawdkcouponabandonAPIRequest, session string) (*promotion.AlibabawdkcouponabandonAPIResponse, error) {
	var resp promotion.AlibabawdkcouponabandonAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
