package btrip

// SeatVo 
type SeatVo struct {

    // 候补价
    
    HoubuPrice   int64 `json:"houbu_price,omitempty" xml:"houbu_price,omitempty"`
    

    // 价格
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 名称
    
    SeatName   string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
    

    // 类型
    
    SeatType   int64 `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
    

    // 下铺价
    
    SleeperPrice   int64 `json:"sleeper_price,omitempty" xml:"sleeper_price,omitempty"`
    

    // 库存
    
    Stock   int64 `json:"stock,omitempty" xml:"stock,omitempty"`
    

}
