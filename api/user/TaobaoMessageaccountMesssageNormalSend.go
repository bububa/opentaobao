package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMessageaccountMesssageNormalSend 下行普通消息
// taobao.messageaccount.messsage.normal.send
//
// 消息号下行单个普通消息
func TaobaoMessageaccountMesssageNormalSend(clt *core.SDKClient, req *user.TaobaoMessageaccountMesssageNormalSendAPIRequest, session string) (*user.TaobaoMessageaccountMesssageNormalSendAPIResponse, error) {
	var resp user.TaobaoMessageaccountMesssageNormalSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
