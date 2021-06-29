package tttm

// StockInfoDTO 
type StockInfoDTO struct {
    // 工厂仓
    FactoryDepot   []ProductInfoDTO `json:"factory_depot,omitempty" xml:"factory_depot>product_info_dto,omitempty"`
    // 电商仓
    ShopDepot   []ProductInfoDTO `json:"shop_depot,omitempty" xml:"shop_depot>product_info_dto,omitempty"`
}
