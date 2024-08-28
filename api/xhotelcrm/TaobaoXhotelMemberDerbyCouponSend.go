package xhotelcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelMemberDerbyCouponSend 发券
// taobao.xhotel.member.derby.coupon.send
//
// 发券
func TaobaoXhotelMemberDerbyCouponSend(ctx context.Context, clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelMemberDerbyCouponSendAPIRequest, resp *xhotelcrm.TaobaoXhotelMemberDerbyCouponSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
