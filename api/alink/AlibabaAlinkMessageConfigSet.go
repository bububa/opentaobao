package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkmessageconfigset 消息提醒开关
// alibaba.alink.message.config.set
//
// 阿里智能消息开关
func Alibabaalinkmessageconfigset(clt *core.SDKClient, req *alink.AlibabaalinkmessageconfigsetAPIRequest, session string) (*alink.AlibabaalinkmessageconfigsetAPIResponse, error) {
	var resp alink.AlibabaalinkmessageconfigsetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
