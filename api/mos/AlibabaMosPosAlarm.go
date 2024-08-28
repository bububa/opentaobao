package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosPosAlarm pos故障报警
// alibaba.mos.pos.alarm
//
// 故障报警
func AlibabaMosPosAlarm(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosPosAlarmAPIRequest, resp *mos.AlibabaMosPosAlarmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
