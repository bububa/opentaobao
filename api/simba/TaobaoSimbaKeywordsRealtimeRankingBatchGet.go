package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsRealtimeRankingBatchGet 获取关键词的新版实时排名
// taobao.simba.keywords.realtime.ranking.batch.get
//
// 根据关键词ID获取关键词的新版实时排名
func TaobaoSimbaKeywordsRealtimeRankingBatchGet(clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest, resp *simba.TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
