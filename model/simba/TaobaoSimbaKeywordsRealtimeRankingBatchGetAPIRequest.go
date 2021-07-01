package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest
获取关键词的新版实时排名 API请求
taobao.simba.keywords.realtime.ranking.batch.get

根据关键词ID获取关键词的新版实时排名 */
type TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIRequest struct {
	model.Params
	// 旺旺名称
	_nick string
	// adgroupId
	_adGroupId int64
	// 关键词列表集合,id用半角逗号分割，一次最多20个
	_bidwordIds []int64
}

// New
