package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveTaskRegister 互动任务开通接口
// taobao.jst.interactive.task.register
//
// 调用互动任务开通接口为小程序开通互动任务
func TaobaoJstInteractiveTaskRegister(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveTaskRegisterAPIRequest, resp *jstinteractive.TaobaoJstInteractiveTaskRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
