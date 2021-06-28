package qt

// ServiceItemProperty 
/* model for simplify = false
type ServiceItemProperty struct {

    // 属性列表
    
    ItemPropertyValues  struct {
        ItemPropertyValues  []ItemPropertyValues `json:"item_property_values,omitempty"`
    } `json:"item_property_values,omitempty"`
    

    // 服务名称
    
    ServiceName   string `json:"service_name,omitempty"`
    

    // 服务收费项名称
    
    ServiceItemName   string `json:"service_item_name,omitempty"`
    

    // 服务收费项代码
    
    ServiceItemCode   string `json:"service_item_code,omitempty"`
    

    // 基础价格
    
    BasicPrice   float64 `json:"basic_price,omitempty"`
    

    // 服务商的nick
    
    Nick   string `json:"nick,omitempty"`
    

    // 质检服务简介
    
    Description   string `json:"description,omitempty"`
    

}
*/

// ServiceItemProperty 
type ServiceItemProperty struct {

    // 属性列表
    ItemPropertyValues   []ItemPropertyValues `json:"item_property_values,omitempty"`

    // 服务名称
    ServiceName   string `json:"service_name,omitempty"`

    // 服务收费项名称
    ServiceItemName   string `json:"service_item_name,omitempty"`

    // 服务收费项代码
    ServiceItemCode   string `json:"service_item_code,omitempty"`

    // 基础价格
    BasicPrice   float64 `json:"basic_price,omitempty"`

    // 服务商的nick
    Nick   string `json:"nick,omitempty"`

    // 质检服务简介
    Description   string `json:"description,omitempty"`

}
