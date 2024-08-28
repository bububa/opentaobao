package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsbykeywordidsGet 根据一个关键词Id列表取得一组关键词
// taobao.simba.keywordsbykeywordids.get
//
// 根据一个关键词Id列表取得一组关键词
func TaobaoSimbaKeywordsbykeywordidsGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsbykeywordidsGetAPIRequest, resp *simba.TaobaoSimbaKeywordsbykeywordidsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
