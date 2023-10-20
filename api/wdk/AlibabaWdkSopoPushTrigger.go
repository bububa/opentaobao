package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdksopopushtrigger 猫超共享库存寄售sopo推送触发
// alibaba.wdk.sopo.push.trigger
//
// 猫超共享库存寄售sopo触发推送给商家
func Alibabawdksopopushtrigger(clt *core.SDKClient, req *wdk.AlibabawdksopopushtriggerAPIRequest, session string) (*wdk.AlibabawdksopopushtriggerAPIResponse, error) {
	var resp wdk.AlibabawdksopopushtriggerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
