package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsvonAdd 创建一批关键词
// taobao.simba.keywordsvon.add
//
// 创建一批关键词
func TaobaoSimbaKeywordsvonAdd(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsvonAddAPIRequest, resp *simba.TaobaoSimbaKeywordsvonAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
