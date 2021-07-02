package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaIworkMcMsgSenddefault 给注册用户发送消息
// alibaba.iwork.mc.msg.senddefault
//
// 给神鲸注册用户发送对应操作结果的消息
func AlibabaIworkMcMsgSenddefault(clt *core.SDKClient, req *campus.AlibabaIworkMcMsgSenddefaultAPIRequest, session string) (*campus.AlibabaIworkMcMsgSenddefaultAPIResponse, error) {
	var resp campus.AlibabaIworkMcMsgSenddefaultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
