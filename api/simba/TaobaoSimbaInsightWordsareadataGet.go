package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaInsightWordsareadataGet 获取关键词按地域进行细分的数据
// taobao.simba.insight.wordsareadata.get
//
// 获取关键词按地域细分的详细数据，目前地域只能细化到省级别，返回的结果中包括市，是为了方便以后扩展，目前结果中市的值等于省。
func TaobaoSimbaInsightWordsareadataGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaInsightWordsareadataGetAPIRequest, resp *simba.TaobaoSimbaInsightWordsareadataGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
