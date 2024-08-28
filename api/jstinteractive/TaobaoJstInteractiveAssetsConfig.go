package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveAssetsConfig 任务素材配置接口
// taobao.jst.interactive.assets.config
//
// 任务素材配置接口
func TaobaoJstInteractiveAssetsConfig(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveAssetsConfigAPIRequest, resp *jstinteractive.TaobaoJstInteractiveAssetsConfigAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
