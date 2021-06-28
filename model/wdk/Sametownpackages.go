package wdk

// Sametownpackages 
/* model for simplify = false
type Sametownpackages struct {

    // 货品列表
    
    SkuDetails  struct {
        Skudetails  []Skudetails `json:"skudetails,omitempty"`
    } `json:"sku_details,omitempty"`
    

    // 扩展信息  json格式
    
    Attribute   string `json:"attribute,omitempty"`
    

    // 是否测试 1:测试 0:非测试
    
    IsTest   string `json:"is_test,omitempty"`
    

    // 同城令牌 即包裹码
    
    TokenCode   string `json:"token_code,omitempty"`
    

    // 仓编码
    
    WarehouseCode   string `json:"warehouse_code,omitempty"`
    

    // 作业单id
    
    WorkOrderId   string `json:"work_order_id,omitempty"`
    

    // 履约单id
    
    FulfillOrderId   string `json:"fulfill_order_id,omitempty"`
    

}
*/

// Sametownpackages 
type Sametownpackages struct {

    // 货品列表
    SkuDetails   []Skudetails `json:"sku_details,omitempty"`

    // 扩展信息  json格式
    Attribute   string `json:"attribute,omitempty"`

    // 是否测试 1:测试 0:非测试
    IsTest   string `json:"is_test,omitempty"`

    // 同城令牌 即包裹码
    TokenCode   string `json:"token_code,omitempty"`

    // 仓编码
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 作业单id
    WorkOrderId   string `json:"work_order_id,omitempty"`

    // 履约单id
    FulfillOrderId   string `json:"fulfill_order_id,omitempty"`

}
