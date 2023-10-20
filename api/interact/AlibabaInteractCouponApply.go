package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractcouponapply 优惠券领取鉴权接口
// alibaba.interact.coupon.apply
//
// 鉴权接口，为coupon.apply接口鉴权
func Alibabainteractcouponapply(clt *core.SDKClient, req *interact.AlibabainteractcouponapplyAPIRequest, session string) (*interact.AlibabainteractcouponapplyAPIResponse, error) {
	var resp interact.AlibabainteractcouponapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
