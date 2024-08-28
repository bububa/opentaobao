package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarKeywordsDelete 销量明星关键词删除
// taobao.simba.salestar.keywords.delete
//
// （新）关键词删除相关接口
func TaobaoSimbaSalestarKeywordsDelete(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSalestarKeywordsDeleteAPIRequest, resp *simba.TaobaoSimbaSalestarKeywordsDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
