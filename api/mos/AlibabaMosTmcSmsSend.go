package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosTmcSmsSend 银泰TMC发送短信
// alibaba.mos.tmc.sms.send
//
// 银泰TMC发送短信
func AlibabaMosTmcSmsSend(clt *core.SDKClient, req *mos.AlibabaMosTmcSmsSendAPIRequest, session string) (*mos.AlibabaMosTmcSmsSendAPIResponse, error) {
	var resp mos.AlibabaMosTmcSmsSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
