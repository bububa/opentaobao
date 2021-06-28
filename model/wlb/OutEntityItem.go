package wlb

// OutEntityItem 
/* model for simplify = false
type OutEntityItem struct {

    // 外部实体类型：<br/>IC_ITEM--ic商品<br/>IC_SKU--ic销售单元
    
    EntityType   string `json:"entity_type,omitempty"`
    

    // entity_type对应的实体类型的id：<br/>当entity_type为IC_ITEM时，entity_id为ic的商品id
    
    EntityId   string `json:"entity_id,omitempty"`
    

}
*/

// OutEntityItem 
type OutEntityItem struct {

    // 外部实体类型：<br/>IC_ITEM--ic商品<br/>IC_SKU--ic销售单元
    EntityType   string `json:"entity_type,omitempty"`

    // entity_type对应的实体类型的id：<br/>当entity_type为IC_ITEM时，entity_id为ic的商品id
    EntityId   string `json:"entity_id,omitempty"`

}
