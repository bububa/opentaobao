package tbk

// TopNInfoDTO 
/* model for simplify = false
type TopNInfoDTO struct {

    // 前N件剩余库存
    
    TopnQuantity   int64 `json:"topn_quantity,omitempty"`
    

    // 前N件初始总库存
    
    TopnTotalCount   int64 `json:"topn_total_count,omitempty"`
    

    // 前N件佣金结束时间
    
    TopnEndTime   string `json:"topn_end_time,omitempty"`
    

    // 前N件佣金开始时间
    
    TopnStartTime   string `json:"topn_start_time,omitempty"`
    

    // 前N件佣金率
    
    TopnRate   string `json:"topn_rate,omitempty"`
    

}
*/

// TopNInfoDTO 
type TopNInfoDTO struct {

    // 前N件剩余库存
    TopnQuantity   int64 `json:"topn_quantity,omitempty"`

    // 前N件初始总库存
    TopnTotalCount   int64 `json:"topn_total_count,omitempty"`

    // 前N件佣金结束时间
    TopnEndTime   string `json:"topn_end_time,omitempty"`

    // 前N件佣金开始时间
    TopnStartTime   string `json:"topn_start_time,omitempty"`

    // 前N件佣金率
    TopnRate   string `json:"topn_rate,omitempty"`

}
