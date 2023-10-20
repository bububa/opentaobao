package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveAssetsQuery 查询可配置任务素材接口
// taobao.jst.interactive.assets.query
//
// 查询可配置任务素材列表，用以配置任务素材
func TaobaoJstInteractiveAssetsQuery(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveAssetsQueryAPIRequest, resp *jstinteractive.TaobaoJstInteractiveAssetsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
