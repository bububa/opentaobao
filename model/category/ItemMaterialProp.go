package category

// ItemMaterialProp 
/* model for simplify = false
type ItemMaterialProp struct {

    // 材质值列表
    
    Materials  struct {
        ItemMateriaValueDO  []ItemMateriaValueDO `json:"item_materia_value_do,omitempty"`
    } `json:"materials,omitempty"`
    

}
*/

// ItemMaterialProp 
type ItemMaterialProp struct {

    // 材质值列表
    Materials   []ItemMateriaValueDO `json:"materials,omitempty"`

}
