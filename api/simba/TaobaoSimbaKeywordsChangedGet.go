package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsChangedGet 分页获取修改过的关键词ID、宝贝id、修改时间
// taobao.simba.keywords.changed.get
//
// 分页获取修改过的关键词ID、宝贝id、修改时间
func TaobaoSimbaKeywordsChangedGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsChangedGetAPIRequest, resp *simba.TaobaoSimbaKeywordsChangedGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
