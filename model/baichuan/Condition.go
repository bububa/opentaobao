package baichuan

// Condition 
type Condition struct {
    // 开始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 查询个数，有数量限制，超过最大值报错
    Limit   int64 `json:"limit,omitempty" xml:"limit,omitempty"`
    // 查询的起始id，如果要连续的分页查询，第n+1页的查询输入是第n页查询结果中最大id+1
    StartId   int64 `json:"start_id,omitempty" xml:"start_id,omitempty"`
    // 商品状态
    ItemStatus   int64 `json:"item_status,omitempty" xml:"item_status,omitempty"`
    // 结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}
