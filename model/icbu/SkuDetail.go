package icbu

// SkuDetail 
/* model for simplify = false
type SkuDetail struct {

    // 商品属性
    
    Attributes  struct {
        ProductAttribute  []ProductAttribute `json:"product_attribute,omitempty"`
    } `json:"attributes,omitempty"`
    

    // 价格，单位是美元，精确到小数点后两位，范围是0.01-9999999.00
    
    Price   string `json:"price,omitempty"`
    

    // 商品的SKU编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 库存
    
    InventoryDtoList  struct {
        InventoryDetail  []InventoryDetail `json:"inventory_detail,omitempty"`
    } `json:"inventory_dto_list,omitempty"`
    

    // SKU id，唯一标识一个SKU
    
    SkuId   int64 `json:"sku_id,omitempty"`
    

}
*/

// SkuDetail 
type SkuDetail struct {

    // 商品属性
    Attributes   []ProductAttribute `json:"attributes,omitempty"`

    // 价格，单位是美元，精确到小数点后两位，范围是0.01-9999999.00
    Price   string `json:"price,omitempty"`

    // 商品的SKU编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 库存
    InventoryDtoList   []InventoryDetail `json:"inventory_dto_list,omitempty"`

    // SKU id，唯一标识一个SKU
    SkuId   int64 `json:"sku_id,omitempty"`

}
