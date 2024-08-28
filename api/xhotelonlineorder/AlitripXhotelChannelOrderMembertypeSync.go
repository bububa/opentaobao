package xhotelonlineorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// AlitripXhotelChannelOrderMembertypeSync 酒店分销渠道会员类型同步
// alitrip.xhotel.channel.order.membertype.sync
//
// 酒店分销渠道会员类型同步
func AlitripXhotelChannelOrderMembertypeSync(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.AlitripXhotelChannelOrderMembertypeSyncAPIRequest, resp *xhotelonlineorder.AlitripXhotelChannelOrderMembertypeSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
