package bus

// TopBusPriceAndStockUpdateRq 
type TopBusPriceAndStockUpdateRq struct {

    // 车次列表
    Numbers   []BusNumberInfoDto `json:"numbers,omitempty"`

}
