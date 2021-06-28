package hotel

// RoomTypeDailyPrice 
/* model for simplify = false
type RoomTypeDailyPrice struct {

    // 床型数据
    
    BedTypeString   string `json:"bed_type_string,omitempty"`
    

    // h5下单页链接
    
    H5BuyUrl   string `json:"h5_buy_url,omitempty"`
    

    // 房型名称
    
    Name   string `json:"name,omitempty"`
    

    // pc下单页链接
    
    PcBuyUrl   string `json:"pc_buy_url,omitempty"`
    

    // 库价具体信息
    
    Rates  struct {
        Rate  []Rate `json:"rate,omitempty"`
    } `json:"rates,omitempty"`
    

    // 卖家房型id
    
    Rid   int64 `json:"rid,omitempty"`
    

    // 标准房型srid
    
    Srid   int64 `json:"srid,omitempty"`
    

    // 床型信息
    
    WindowType   string `json:"window_type,omitempty"`
    

}
*/

// RoomTypeDailyPrice 
type RoomTypeDailyPrice struct {

    // 床型数据
    BedTypeString   string `json:"bed_type_string,omitempty"`

    // h5下单页链接
    H5BuyUrl   string `json:"h5_buy_url,omitempty"`

    // 房型名称
    Name   string `json:"name,omitempty"`

    // pc下单页链接
    PcBuyUrl   string `json:"pc_buy_url,omitempty"`

    // 库价具体信息
    Rates   []Rate `json:"rates,omitempty"`

    // 卖家房型id
    Rid   int64 `json:"rid,omitempty"`

    // 标准房型srid
    Srid   int64 `json:"srid,omitempty"`

    // 床型信息
    WindowType   string `json:"window_type,omitempty"`

}
