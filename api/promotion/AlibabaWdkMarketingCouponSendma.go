package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkmarketingcouponsendma 发放匿名码
// alibaba.wdk.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
func Alibabawdkmarketingcouponsendma(clt *core.SDKClient, req *promotion.AlibabawdkmarketingcouponsendmaAPIRequest, session string) (*promotion.AlibabawdkmarketingcouponsendmaAPIResponse, error) {
	var resp promotion.AlibabawdkmarketingcouponsendmaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
