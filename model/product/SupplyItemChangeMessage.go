package product

// SupplyItemChangeMessage 
/* model for simplify = false
type SupplyItemChangeMessage struct {

    // 货号列表
    
    ProductCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"product_codes,omitempty"`
    

    // 供应商门店，最大20个
    
    SupplyStoreId   string `json:"supply_store_id,omitempty"`
    

}
*/

// SupplyItemChangeMessage 
type SupplyItemChangeMessage struct {

    // 货号列表
    ProductCodes   []string `json:"product_codes,omitempty"`

    // 供应商门店，最大20个
    SupplyStoreId   string `json:"supply_store_id,omitempty"`

}
