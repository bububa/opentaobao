package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskcreate 创建轻任务
// taobao.qianniu.task.create
//
// 发起一个轻任务，分配给多个执行者，并发送消息提醒，由任务发起者调用
func Taobaoqianniutaskcreate(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskcreateAPIRequest, session string) (*qianniu.TaobaoqianniutaskcreateAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
