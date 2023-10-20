package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcoupontemplateterminate 优惠券模版终止
// alibaba.wdk.coupon.template.terminate
//
// 优惠券模版终止
func Alibabawdkcoupontemplateterminate(clt *core.SDKClient, req *promotion.AlibabawdkcoupontemplateterminateAPIRequest, session string) (*promotion.AlibabawdkcoupontemplateterminateAPIResponse, error) {
	var resp promotion.AlibabawdkcoupontemplateterminateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
