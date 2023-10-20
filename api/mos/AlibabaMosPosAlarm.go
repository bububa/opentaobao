package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosposalarm pos故障报警
// alibaba.mos.pos.alarm
//
// 故障报警
func Alibabamosposalarm(clt *core.SDKClient, req *mos.AlibabamosposalarmAPIRequest, session string) (*mos.AlibabamosposalarmAPIResponse, error) {
	var resp mos.AlibabamosposalarmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
