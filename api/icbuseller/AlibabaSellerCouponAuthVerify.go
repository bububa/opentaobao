package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// Alibabasellercouponauthverify 优惠券校验
// alibaba.seller.coupon.auth.verify
//
// 优惠券校验
func Alibabasellercouponauthverify(clt *core.SDKClient, req *icbuseller.AlibabasellercouponauthverifyAPIRequest, session string) (*icbuseller.AlibabasellercouponauthverifyAPIResponse, error) {
	var resp icbuseller.AlibabasellercouponauthverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
