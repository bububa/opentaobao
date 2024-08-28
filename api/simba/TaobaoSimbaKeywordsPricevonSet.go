package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsPricevonSet 设置一批关键词的信息
// taobao.simba.keywords.pricevon.set
//
// 设置一批关键词的信息，包含无线出价、计算机出价和关键词匹配方式
func TaobaoSimbaKeywordsPricevonSet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsPricevonSetAPIRequest, resp *simba.TaobaoSimbaKeywordsPricevonSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
