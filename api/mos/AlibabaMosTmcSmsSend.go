package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamostmcsmssend 银泰TMC发送短信
// alibaba.mos.tmc.sms.send
//
// 银泰TMC发送短信
func Alibabamostmcsmssend(clt *core.SDKClient, req *mos.AlibabamostmcsmssendAPIRequest, session string) (*mos.AlibabamostmcsmssendAPIResponse, error) {
	var resp mos.AlibabamostmcsmssendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
