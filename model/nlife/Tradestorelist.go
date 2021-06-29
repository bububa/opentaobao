package nlife

// Tradestorelist 
type Tradestorelist struct {
    // 门店采购单号
    TradeNo   string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
    // 供应商ID
    SupplierId   int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 门店ID
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
