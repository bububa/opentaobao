package paimai

// ItemMaterialProp 
type ItemMaterialProp struct {
    // 材质值列表
    Materials   []ItemMateriaValueDo `json:"materials,omitempty" xml:"materials>item_materia_value_do,omitempty"`
}
