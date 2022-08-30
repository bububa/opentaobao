package aedropshiper

// TrafficOrderResultDto 结构体
type TrafficOrderResultDto struct {
	// orders object list
	Orders []TrafficOrderEffectDto `json:"orders,omitempty" xml:"orders>traffic_order_effect_dto,omitempty"`
	// max query index start value: if not passed, You can only check the first page
	MaxQueryIndexId string `json:"max_query_index_id,omitempty" xml:"max_query_index_id,omitempty"`
	// min query index start value: if not passed, You can only check the first page
	MinQueryIndexId string `json:"min_query_index_id,omitempty" xml:"min_query_index_id,omitempty"`
	// current record count
	CurrentRecordCount int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
	// current page number
	CurrentPageNo int64 `json:"current_page_no,omitempty" xml:"current_page_no,omitempty"`
}
