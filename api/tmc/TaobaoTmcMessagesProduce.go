package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcMessagesProduce 批量发送消息
// taobao.tmc.messages.produce
//
// 批量发送消息
func TaobaoTmcMessagesProduce(clt *core.SDKClient, req *tmc.TaobaoTmcMessagesProduceAPIRequest, resp *tmc.TaobaoTmcMessagesProduceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
