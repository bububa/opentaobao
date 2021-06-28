package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
pos故障报警 
alibaba.mos.pos.alarm

故障报警
*/
func AlibabaMosPosAlarm(clt *core.SDKClient, req *mos.AlibabaMosPosAlarmRequest, session string) (*mos.AlibabaMosPosAlarmAPIResponse, error) {
    var resp mos.AlibabaMosPosAlarmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
