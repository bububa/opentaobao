package simba

// TaobaoSimbaKeywordsRealtimeRankingBatchGetResultDto 
type TaobaoSimbaKeywordsRealtimeRankingBatchGetResultDto struct {

    // 返回结果，结果是json结构，说明： pc_rank=-2:创意未投放,-1:计划未投放,0:首页左侧位置,1:首页右侧第1,2:首页右侧第2,3:首页右侧第3,4:首页（非前三）,5:第2页,6:第3页,7:第4页,8:第5页,9:5页以后 mobile_rank=-2:创意未投放,-1:计划未投放,0:移动首条,1:移动前三,3:移动4~6条,6:移动7~10条,10:移动11~15条,11=移动16~20条,12=20条以后
    
    RealtimeRankList   []TaobaoSimbaKeywordsRealtimeRankingBatchGetResult `json:"realtime_rank_list,omitempty" xml:"realtime_rank_list,omitempty"`
    

    // 错误码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误消息
    
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    

    // 错误对应的消息Key
    
    Key   string `json:"key,omitempty" xml:"key,omitempty"`
    

}
