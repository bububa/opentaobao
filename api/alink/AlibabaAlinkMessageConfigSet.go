package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkMessageConfigSet 消息提醒开关
// alibaba.alink.message.config.set
//
// 阿里智能消息开关
func AlibabaAlinkMessageConfigSet(clt *core.SDKClient, req *alink.AlibabaAlinkMessageConfigSetAPIRequest, resp *alink.AlibabaAlinkMessageConfigSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
