package scbp

// BidStrategyDto 
type BidStrategyDto struct {

    // 主键
    Id   string `json:"id,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 目标排名，前N名
    Topn   int64 `json:"topn,omitempty"`

    // 出价间隔分钟为单位的数值型
    Duration   int64 `json:"duration,omitempty"`

    // 产品
    ProductId   int64 `json:"product_id,omitempty"`

    // 溢价比例
    Discount   int64 `json:"discount,omitempty"`

}
