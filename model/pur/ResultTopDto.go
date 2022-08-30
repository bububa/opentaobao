package pur

// ResultTopDto 结构体
type ResultTopDto struct {
	// 行映射关系
	IdRelationList []RelationTopDto `json:"id_relation_list,omitempty" xml:"id_relation_list>relation_top_dto,omitempty"`
	// 拓展字段
	ExtStr string `json:"ext_str,omitempty" xml:"ext_str,omitempty"`
	// 发货单号
	Number string `json:"number,omitempty" xml:"number,omitempty"`
	// 发货单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
