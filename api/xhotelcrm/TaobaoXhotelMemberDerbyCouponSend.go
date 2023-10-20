package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelMemberDerbyCouponSend 发券
// taobao.xhotel.member.derby.coupon.send
//
// 发券
func TaobaoXhotelMemberDerbyCouponSend(clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelMemberDerbyCouponSendAPIRequest, resp *xhotelcrm.TaobaoXhotelMemberDerbyCouponSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
