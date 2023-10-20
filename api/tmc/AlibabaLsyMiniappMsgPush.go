package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Alibabalsyminiappmsgpush 零售云小程序消息推送
// alibaba.lsy.miniapp.msg.push
//
// 零售云小程序消息推送，推送消息至零售云（喵零等）
func Alibabalsyminiappmsgpush(clt *core.SDKClient, req *tmc.AlibabalsyminiappmsgpushAPIRequest, session string) (*tmc.AlibabalsyminiappmsgpushAPIResponse, error) {
	var resp tmc.AlibabalsyminiappmsgpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
