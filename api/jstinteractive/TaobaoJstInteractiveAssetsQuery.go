package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveAssetsQuery 查询可配置任务素材接口
// taobao.jst.interactive.assets.query
//
// 查询可配置任务素材列表，用以配置任务素材
func TaobaoJstInteractiveAssetsQuery(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveAssetsQueryAPIRequest, resp *jstinteractive.TaobaoJstInteractiveAssetsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
