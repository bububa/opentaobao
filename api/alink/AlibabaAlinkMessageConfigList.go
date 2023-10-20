package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkmessageconfiglist 查询消息开关配置列表
// alibaba.alink.message.config.list
//
// 阿里智能获取消息开关配置列表
func Alibabaalinkmessageconfiglist(clt *core.SDKClient, req *alink.AlibabaalinkmessageconfiglistAPIRequest, session string) (*alink.AlibabaalinkmessageconfiglistAPIResponse, error) {
	var resp alink.AlibabaalinkmessageconfiglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
