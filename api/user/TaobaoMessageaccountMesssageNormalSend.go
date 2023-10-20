package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMessageaccountMesssageNormalSend 下行普通消息
// taobao.messageaccount.messsage.normal.send
//
// 消息号下行单个普通消息
func TaobaoMessageaccountMesssageNormalSend(clt *core.SDKClient, req *user.TaobaoMessageaccountMesssageNormalSendAPIRequest, resp *user.TaobaoMessageaccountMesssageNormalSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
