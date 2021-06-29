package wlbimports

// TranStoreResult 
type TranStoreResult struct {
    // 中转仓名字
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 中转仓地址
    StoreAddress   string `json:"store_address,omitempty" xml:"store_address,omitempty"`
    // 中转仓代码
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}
