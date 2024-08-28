package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripUppcMemberGain 航司权益数据回流
// alitrip.uppc.member.gain
//
// 航司权益数据回流
func AlitripUppcMemberGain(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripUppcMemberGainAPIRequest, resp *flightuppc.AlitripUppcMemberGainAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
