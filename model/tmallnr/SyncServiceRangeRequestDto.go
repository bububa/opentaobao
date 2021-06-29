package tmallnr

// SyncServiceRangeRequestDTO 
type SyncServiceRangeRequestDTO struct {
    // 围栏信息
    Points   []Point `json:"points,omitempty" xml:"points>point,omitempty"`
    // 扩展字段
    Properties   string `json:"properties,omitempty" xml:"properties,omitempty"`
    // 门店id
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 类型 -1 为删除,其余为添加，如果记录存在则做更新
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
}
