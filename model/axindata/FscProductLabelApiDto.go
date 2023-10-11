package axindata

// FscProductLabelApiDto 结构体
type FscProductLabelApiDto struct {
	// 线路主题ID
	ProductLabelId string `json:"product_label_id,omitempty" xml:"product_label_id,omitempty"`
	// 线路主题名称
	ProductLabelName string `json:"product_label_name,omitempty" xml:"product_label_name,omitempty"`
	// 父级线路主题ID
	ParentLabelId string `json:"parent_label_id,omitempty" xml:"parent_label_id,omitempty"`
	// 根节点线路主题ID
	RootLabelId string `json:"root_label_id,omitempty" xml:"root_label_id,omitempty"`
	// 主题层级
	ProductLabelLevel int64 `json:"product_label_level,omitempty" xml:"product_label_level,omitempty"`
}
