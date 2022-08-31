package alihouse

// RelationBindingDto 结构体
type RelationBindingDto struct {
	// 货下挂的其他品列表 最大列表长度：100
	Extend []GoodsRelationDto `json:"extend,omitempty" xml:"extend>goods_relation_dto,omitempty"`
	// 货下挂的其他货列表 最大列表长度：100
	RelationCargos []GoodsRelationDto `json:"relation_cargos,omitempty" xml:"relation_cargos>goods_relation_dto,omitempty"`
	// 外部货id（外部唯一识别码）
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 外部私域楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}
