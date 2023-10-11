package alicom

// GroupRequest 结构体
type GroupRequest struct {
	// SPU ID列表
	SpuIdList []int64 `json:"spu_id_list,omitempty" xml:"spu_id_list>int64,omitempty"`
	// 是否主接口
	Main bool `json:"main,omitempty" xml:"main,omitempty"`
}
