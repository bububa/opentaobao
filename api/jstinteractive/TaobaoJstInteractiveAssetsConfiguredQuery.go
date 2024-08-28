package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveAssetsConfiguredQuery 查询已配置的任务素材列表接口
// taobao.jst.interactive.assets.configured.query
//
// 查询已配置任务素材列表
func TaobaoJstInteractiveAssetsConfiguredQuery(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest, resp *jstinteractive.TaobaoJstInteractiveAssetsConfiguredQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
