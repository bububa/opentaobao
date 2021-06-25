package product

// SupplyItemChangeMessage 
type SupplyItemChangeMessage struct {

    // 货号列表
    ProductCodes   []String `json:"product_codes,omitempty"`

    // 供应商门店，最大20个
    SupplyStoreId   string `json:"supply_store_id,omitempty"`

}
