package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcouponskuremove 优惠券商品删除
// alibaba.wdk.coupon.sku.remove
//
// 优惠券商品删除
func Alibabawdkcouponskuremove(clt *core.SDKClient, req *promotion.AlibabawdkcouponskuremoveAPIRequest, session string) (*promotion.AlibabawdkcouponskuremoveAPIResponse, error) {
	var resp promotion.AlibabawdkcouponskuremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
