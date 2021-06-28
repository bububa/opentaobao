package hotel

// HotelInfoParam 
/* model for simplify = false
type HotelInfoParam struct {

    // 城市code
    
    CityCode   int64 `json:"city_code,omitempty"`
    

    // 上次请求的分页数据中最后一个shid
    
    LastShid   int64 `json:"last_shid,omitempty"`
    

    // 每页数据量
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // pid
    
    Pid   string `json:"pid,omitempty"`
    

    // 单个请求的shid
    
    Shid   int64 `json:"shid,omitempty"`
    

    // 经过混淆的userId
    
    OpenId   string `json:"open_id,omitempty"`
    

}
*/

// HotelInfoParam 
type HotelInfoParam struct {

    // 城市code
    CityCode   int64 `json:"city_code,omitempty"`

    // 上次请求的分页数据中最后一个shid
    LastShid   int64 `json:"last_shid,omitempty"`

    // 每页数据量
    PageSize   int64 `json:"page_size,omitempty"`

    // pid
    Pid   string `json:"pid,omitempty"`

    // 单个请求的shid
    Shid   int64 `json:"shid,omitempty"`

    // 经过混淆的userId
    OpenId   string `json:"open_id,omitempty"`

}
