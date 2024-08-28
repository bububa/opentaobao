package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordUpdate （新）关键词更新相关接口
// taobao.simba.keyword.update
//
// （新）关键词更新相关接口
func TaobaoSimbaKeywordUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordUpdateAPIRequest, resp *simba.TaobaoSimbaKeywordUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
