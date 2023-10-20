package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosPosAlarm pos故障报警
// alibaba.mos.pos.alarm
//
// 故障报警
func AlibabaMosPosAlarm(clt *core.SDKClient, req *mos.AlibabaMosPosAlarmAPIRequest, resp *mos.AlibabaMosPosAlarmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
