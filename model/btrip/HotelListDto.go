package btrip

// HotelListDto 
type HotelListDto struct {
    // 最低价，单位分
    LowPrice   int64 `json:"low_price,omitempty" xml:"low_price,omitempty"`
    // 酒店名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 酒店标准ID
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
    // 供应商code
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // 供应商名称
    SupplierName   string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
}
