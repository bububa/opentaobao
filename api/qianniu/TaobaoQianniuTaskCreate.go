package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskCreate 创建轻任务
// taobao.qianniu.task.create
//
// 发起一个轻任务，分配给多个执行者，并发送消息提醒，由任务发起者调用
func TaobaoQianniuTaskCreate(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskCreateAPIRequest, resp *qianniu.TaobaoQianniuTaskCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
