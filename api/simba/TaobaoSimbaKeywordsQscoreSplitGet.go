package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsQscoreSplitGet 新质量分服务
// taobao.simba.keywords.qscore.split.get
//
// 获取关键词新的质量分
func TaobaoSimbaKeywordsQscoreSplitGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsQscoreSplitGetAPIRequest, resp *simba.TaobaoSimbaKeywordsQscoreSplitGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
