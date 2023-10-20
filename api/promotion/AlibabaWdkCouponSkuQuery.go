package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcouponskuquery 优惠券商品查询
// alibaba.wdk.coupon.sku.query
//
// 优惠券商品查询
func Alibabawdkcouponskuquery(clt *core.SDKClient, req *promotion.AlibabawdkcouponskuqueryAPIRequest, session string) (*promotion.AlibabawdkcouponskuqueryAPIResponse, error) {
	var resp promotion.AlibabawdkcouponskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
