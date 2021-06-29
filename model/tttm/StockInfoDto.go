package tttm

// StockInfoDto 
type StockInfoDto struct {

    // 工厂仓
    
    FactoryDepot   []ProductInfoDto `json:"factory_depot,omitempty" xml:"factory_depot,omitempty"`
    

    // 电商仓
    
    ShopDepot   []ProductInfoDto `json:"shop_depot,omitempty" xml:"shop_depot,omitempty"`
    

}
