package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaIworkMcMsgSendmobile 发送消息给手机用户
// alibaba.iwork.mc.msg.sendmobile
//
// 给手机用户发送对应操作结果的消息
func AlibabaIworkMcMsgSendmobile(clt *core.SDKClient, req *campus.AlibabaIworkMcMsgSendmobileAPIRequest, resp *campus.AlibabaIworkMcMsgSendmobileAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
