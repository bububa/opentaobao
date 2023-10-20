package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveActivityUpdate 互动任务活动修改接口
// taobao.jst.interactive.activity.update
//
// 互动任务活动修改接口
func TaobaoJstInteractiveActivityUpdate(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveActivityUpdateAPIRequest, resp *jstinteractive.TaobaoJstInteractiveActivityUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
