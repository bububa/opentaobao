package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaIworkMcMsgSenddefault 给注册用户发送消息
// alibaba.iwork.mc.msg.senddefault
//
// 给神鲸注册用户发送对应操作结果的消息
func AlibabaIworkMcMsgSenddefault(clt *core.SDKClient, req *campus.AlibabaIworkMcMsgSenddefaultAPIRequest, resp *campus.AlibabaIworkMcMsgSenddefaultAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
