package icbu

// LevelAttributeValueRequest 
type LevelAttributeValueRequest struct {

    // 必填；要查询的属性值所属发布类目
    CatId   int64 `json:"cat_id,omitempty"`

    // 类目属性id，放到数组第一个位置
    AttrId   []Number `json:"attr_id,omitempty"`

    // 属性值id, 不同取值范围时的查询策略如下:  <=0：列出当前类目属性的所有属性值  >0：指定当前类目属性的某一个属性值，列出该属性值下的子属性和该子属性的所有属性值
    ValueId   int64 `json:"value_id,omitempty"`

}