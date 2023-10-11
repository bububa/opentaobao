package simba

// PreAddItemCreativeVo 结构体
type PreAddItemCreativeVo struct {
	// 单品创意素材
	Material *ItemMaterialVo `json:"material,omitempty" xml:"material,omitempty"`
	// 是否有长素材
	HaveLongSucai bool `json:"have_long_sucai,omitempty" xml:"have_long_sucai,omitempty"`
}
