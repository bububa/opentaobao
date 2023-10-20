package msgamp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/msgamp"
)

// TaobaoMessageSend 消息发送
// taobao.message.send
//
// 消息发送接口
func TaobaoMessageSend(clt *core.SDKClient, req *msgamp.TaobaoMessageSendAPIRequest, session string) (*msgamp.TaobaoMessageSendAPIResponse, error) {
	var resp msgamp.TaobaoMessageSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
