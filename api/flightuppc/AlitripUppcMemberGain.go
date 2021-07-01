package flightuppc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flightuppc"
)

/* 
航司权益数据回流 
alitrip.uppc.member.gain

航司权益数据回流
*/
func AlitripUppcMemberGain(clt *core.SDKClient, req *flightuppc.AlitripUppcMemberGainAPIRequest, session string) (*flightuppc.AlitripUppcMemberGainAPIResponse, error) {
    var resp flightuppc.AlitripUppcMemberGainAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
