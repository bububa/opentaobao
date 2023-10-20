package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveActivityQuery 互动任务活动查询接口
// taobao.jst.interactive.activity.query
//
// 互动任务活动查询接口
func TaobaoJstInteractiveActivityQuery(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveActivityQueryAPIRequest, session string) (*jstinteractive.TaobaoJstInteractiveActivityQueryAPIResponse, error) {
	var resp jstinteractive.TaobaoJstInteractiveActivityQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
