package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabawdkcoupontemplateupdate 优惠券模版修改
// alibaba.wdk.coupon.template.update
//
// 优惠券模版修改
func Alibabawdkcoupontemplateupdate(clt *core.SDKClient, req *promotion.AlibabawdkcoupontemplateupdateAPIRequest, session string) (*promotion.AlibabawdkcoupontemplateupdateAPIResponse, error) {
	var resp promotion.AlibabawdkcoupontemplateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
