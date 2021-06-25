package util

// PosInfoDto 
type PosInfoDto struct {

    // 是否支持小数
    Support4Decimal   string `json:"support4_decimal,omitempty"`

    // 专柜号
    CounterNo   string `json:"counter_no,omitempty"`

    // 门店号
    StoreNo   string `json:"store_no,omitempty"`

}
