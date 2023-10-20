package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcoupontemplatequery 优惠券模版查询
// alibaba.wdk.coupon.template.query
//
// 优惠券模版查询
func Alibabawdkcoupontemplatequery(clt *core.SDKClient, req *promotion.AlibabawdkcoupontemplatequeryAPIRequest, session string) (*promotion.AlibabawdkcoupontemplatequeryAPIResponse, error) {
	var resp promotion.AlibabawdkcoupontemplatequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
