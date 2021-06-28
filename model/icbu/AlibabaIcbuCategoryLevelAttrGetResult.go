package icbu

// AlibabaIcbuCategoryLevelAttrGetResult 
/* model for simplify = false
type AlibabaIcbuCategoryLevelAttrGetResult struct {

    // List<Map<String,Object>>  列表中每个元素的key-value说明如下:  id: 值id  name：值名称  leaf: 此key存在且为true代表当前节点下已无下层属性,这种情况下前端不需再在当前节点上提供弹出下级菜单之类的操作
    
    Values   string `json:"values,omitempty"`
    

    // propertyId对应的属性中文名
    
    PropertyCnName   string `json:"property_cn_name,omitempty"`
    

    // propertyId对应的属性英文名
    
    PropertyEnName   string `json:"property_en_name,omitempty"`
    

    // 返回值所在的属性id，如入参valueId为0，则与入参的attrId一致，否则为所选属性值的下层属性id
    
    PropertyId   int64 `json:"property_id,omitempty"`
    

}
*/

// AlibabaIcbuCategoryLevelAttrGetResult 
type AlibabaIcbuCategoryLevelAttrGetResult struct {

    // List<Map<String,Object>>  列表中每个元素的key-value说明如下:  id: 值id  name：值名称  leaf: 此key存在且为true代表当前节点下已无下层属性,这种情况下前端不需再在当前节点上提供弹出下级菜单之类的操作
    Values   string `json:"values,omitempty"`

    // propertyId对应的属性中文名
    PropertyCnName   string `json:"property_cn_name,omitempty"`

    // propertyId对应的属性英文名
    PropertyEnName   string `json:"property_en_name,omitempty"`

    // 返回值所在的属性id，如入参valueId为0，则与入参的attrId一致，否则为所选属性值的下层属性id
    PropertyId   int64 `json:"property_id,omitempty"`

}
