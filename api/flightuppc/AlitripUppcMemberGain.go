package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripUppcMemberGain 航司权益数据回流
// alitrip.uppc.member.gain
//
// 航司权益数据回流
func AlitripUppcMemberGain(clt *core.SDKClient, req *flightuppc.AlitripUppcMemberGainAPIRequest, resp *flightuppc.AlitripUppcMemberGainAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
