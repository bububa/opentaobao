package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcoupontemplatecreate 优惠券模版创建
// alibaba.wdk.coupon.template.create
//
// 开放给外部商家创建优惠券模版
func Alibabawdkcoupontemplatecreate(clt *core.SDKClient, req *promotion.AlibabawdkcoupontemplatecreateAPIRequest, session string) (*promotion.AlibabawdkcoupontemplatecreateAPIResponse, error) {
	var resp promotion.AlibabawdkcoupontemplatecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
