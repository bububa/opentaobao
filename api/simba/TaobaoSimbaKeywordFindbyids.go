package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordFindbyids （新）根据一堆关键词ids获取关键词
// taobao.simba.keyword.findbyids
//
// 根据一个关键词Id列表取得一组关键词
func TaobaoSimbaKeywordFindbyids(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordFindbyidsAPIRequest, resp *simba.TaobaoSimbaKeywordFindbyidsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
