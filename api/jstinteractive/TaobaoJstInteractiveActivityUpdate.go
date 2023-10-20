package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveActivityUpdate 互动任务活动修改接口
// taobao.jst.interactive.activity.update
//
// 互动任务活动修改接口
func TaobaoJstInteractiveActivityUpdate(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveActivityUpdateAPIRequest, session string) (*jstinteractive.TaobaoJstInteractiveActivityUpdateAPIResponse, error) {
	var resp jstinteractive.TaobaoJstInteractiveActivityUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
