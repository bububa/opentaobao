package simba

import (
	"sync"
)

// TaobaoSimbaKeywordsRealtimeRankingBatchGetResult 结构体
type TaobaoSimbaKeywordsRealtimeRankingBatchGetResult struct {
	// 关键词id
	Bidwordid string `json:"bidwordid,omitempty" xml:"bidwordid,omitempty"`
	// 状态码stat(0:正常;1:非正常)
	Stat string `json:"stat,omitempty" xml:"stat,omitempty"`
	// PC的排名:(0:首页左侧位置;1:首页右侧第1;2:首页右侧第2,3:首页右侧第3,4:首页非前三,5:第2页,6:第3页,7:第4页,8:第5页,9:5页以后)
	PcRank string `json:"pc_rank,omitempty" xml:"pc_rank,omitempty"`
	// 移动的排名(0:移动首条,1:移动前三,3:移动4~6条,6:移动7~10条,10:移动11~15条,11:移动16~20条,12:20条以后)
	MobileRank string `json:"mobile_rank,omitempty" xml:"mobile_rank,omitempty"`
}

var poolTaobaoSimbaKeywordsRealtimeRankingBatchGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsRealtimeRankingBatchGetResult)
	},
}

// GetTaobaoSimbaKeywordsRealtimeRankingBatchGetResult() 从对象池中获取TaobaoSimbaKeywordsRealtimeRankingBatchGetResult
func GetTaobaoSimbaKeywordsRealtimeRankingBatchGetResult() *TaobaoSimbaKeywordsRealtimeRankingBatchGetResult {
	return poolTaobaoSimbaKeywordsRealtimeRankingBatchGetResult.Get().(*TaobaoSimbaKeywordsRealtimeRankingBatchGetResult)
}

// ReleaseTaobaoSimbaKeywordsRealtimeRankingBatchGetResult 释放TaobaoSimbaKeywordsRealtimeRankingBatchGetResult
func ReleaseTaobaoSimbaKeywordsRealtimeRankingBatchGetResult(v *TaobaoSimbaKeywordsRealtimeRankingBatchGetResult) {
	v.Bidwordid = ""
	v.Stat = ""
	v.PcRank = ""
	v.MobileRank = ""
	poolTaobaoSimbaKeywordsRealtimeRankingBatchGetResult.Put(v)
}
