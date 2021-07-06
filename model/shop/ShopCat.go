package shop

// ShopCat 结构体
type ShopCat struct {
	// 类目名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 类目编号
	Cid int64 `json:"cid,omitempty" xml:"cid,omitempty"`
	// 父类目编号，注：此类目指前台类目，值等于0：表示此类目为一级类目，值不等于0：表示此类目有父类目
	ParentCid int64 `json:"parent_cid,omitempty" xml:"parent_cid,omitempty"`
	// 该类目是否为父类目。即：该类目是否还有子类目
	IsParent bool `json:"is_parent,omitempty" xml:"is_parent,omitempty"`
}
