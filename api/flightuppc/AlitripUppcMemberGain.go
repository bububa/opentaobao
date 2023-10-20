package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripuppcmembergain 航司权益数据回流
// alitrip.uppc.member.gain
//
// 航司权益数据回流
func Alitripuppcmembergain(clt *core.SDKClient, req *flightuppc.AlitripuppcmembergainAPIRequest, session string) (*flightuppc.AlitripuppcmembergainAPIResponse, error) {
	var resp flightuppc.AlitripuppcmembergainAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
