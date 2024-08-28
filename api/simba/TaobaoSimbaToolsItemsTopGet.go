package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaToolsItemsTopGet 取得一个关键词的推广组排名列表
// taobao.simba.tools.items.top.get
//
// 取得一个关键词的推广组排名列表
func TaobaoSimbaToolsItemsTopGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaToolsItemsTopGetAPIRequest, resp *simba.TaobaoSimbaToolsItemsTopGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
