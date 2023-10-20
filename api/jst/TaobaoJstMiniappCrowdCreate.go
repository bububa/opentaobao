package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstMiniappCrowdCreate 小程序活动创建
// taobao.jst.miniapp.crowd.create
//
// 小程序活动创建
func TaobaoJstMiniappCrowdCreate(clt *core.SDKClient, req *jst.TaobaoJstMiniappCrowdCreateAPIRequest, resp *jst.TaobaoJstMiniappCrowdCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
