package simba

// CreativeChildrenVo 结构体
type CreativeChildrenVo struct {
	// 素材集合
	MaterialList []ItemMaterialVo `json:"material_list,omitempty" xml:"material_list>item_material_vo,omitempty"`
}
