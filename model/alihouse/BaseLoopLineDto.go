package alihouse

// BaseLoopLineDto 
type BaseLoopLineDto struct {
    // 环线名称
    LoopLineName   string `json:"loop_line_name,omitempty" xml:"loop_line_name,omitempty"`
    // 环线电子围栏
    LoopLineFence   string `json:"loop_line_fence,omitempty" xml:"loop_line_fence,omitempty"`
    // 城市ID
    CityId   int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
    // 序号
    Number   int64 `json:"number,omitempty" xml:"number,omitempty"`
    // 环线ID
    LoopLineId   int64 `json:"loop_line_id,omitempty" xml:"loop_line_id,omitempty"`
    // 外部ID - 唯一
    OuterLineId   string `json:"outer_line_id,omitempty" xml:"outer_line_id,omitempty"`
    // 是否删除 1 是 0 否
    IsDeleted   string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}
