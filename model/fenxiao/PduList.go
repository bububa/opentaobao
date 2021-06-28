package fenxiao

// PduList 
/* model for simplify = false
type PduList struct {

    // 产品ID
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // 分销商ID
    
    DistributorId   int64 `json:"distributor_id,omitempty"`
    

    // 产品销售属性值
    
    SkuProperties   string `json:"sku_properties,omitempty"`
    

    // 产品代销配额库存
    
    QuantityAgent   int64 `json:"quantity_agent,omitempty"`
    

    // 分销商用户名
    
    DistributorName   string `json:"distributor_name,omitempty"`
    

}
*/

// PduList 
type PduList struct {

    // 产品ID
    ProductId   int64 `json:"product_id,omitempty"`

    // 分销商ID
    DistributorId   int64 `json:"distributor_id,omitempty"`

    // 产品销售属性值
    SkuProperties   string `json:"sku_properties,omitempty"`

    // 产品代销配额库存
    QuantityAgent   int64 `json:"quantity_agent,omitempty"`

    // 分销商用户名
    DistributorName   string `json:"distributor_name,omitempty"`

}
