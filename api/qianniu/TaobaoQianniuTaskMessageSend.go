package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskMessageSend 发送任务提醒消息
// taobao.qianniu.task.message.send
//
// 如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。消息会以任务消息的形式发给客户端。
func TaobaoQianniuTaskMessageSend(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskMessageSendAPIRequest, resp *qianniu.TaobaoQianniuTaskMessageSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
