package btrip

// HotelInfoListRs 
type HotelInfoListRs struct {

    // 基础酒店数据列表
    
    Hotels   []HotelDto `json:"hotels,omitempty" xml:"hotels,omitempty"`
    

    // 酒店数量
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

}
