package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcouponskuadd 优惠券商品增加
// alibaba.wdk.coupon.sku.add
//
// 优惠券商品增加
func Alibabawdkcouponskuadd(clt *core.SDKClient, req *promotion.AlibabawdkcouponskuaddAPIRequest, session string) (*promotion.AlibabawdkcouponskuaddAPIResponse, error) {
	var resp promotion.AlibabawdkcouponskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
