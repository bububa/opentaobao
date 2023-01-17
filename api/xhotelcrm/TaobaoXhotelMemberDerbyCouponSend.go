package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelMemberDerbyCouponSend 发券
// taobao.xhotel.member.derby.coupon.send
//
// 发券
func TaobaoXhotelMemberDerbyCouponSend(clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelMemberDerbyCouponSendAPIRequest, session string) (*xhotelcrm.TaobaoXhotelMemberDerbyCouponSendAPIResponse, error) {
	var resp xhotelcrm.TaobaoXhotelMemberDerbyCouponSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
