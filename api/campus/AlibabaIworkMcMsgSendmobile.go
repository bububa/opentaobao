package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

/* AlibabaIworkMcMsgSendmobile
发送消息给手机用户
alibaba.iwork.mc.msg.sendmobile

给手机用户发送对应操作结果的消息 */
func AlibabaIworkMcMsgSendmobile(clt *core.SDKClient, req *campus.AlibabaIworkMcMsgSendmobileAPIRequest, session string) (*campus.AlibabaIworkMcMsgSendmobileAPIResponse, error) {
	var resp campus.AlibabaIworkMcMsgSendmobileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
