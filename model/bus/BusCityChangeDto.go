package bus

// BusCityChangeDto 
/* model for simplify = false
type BusCityChangeDto struct {

    // 是否开通标示，0：已开通 1：未开通
    
    TypeNo   int64 `json:"type_no,omitempty"`
    

    // 城市拼音
    
    CityFullPinyin   string `json:"city_full_pinyin,omitempty"`
    

    // 是否可售标示，0：暂停售票 1：可售
    
    Status   int64 `json:"status,omitempty"`
    

    // 城市名
    
    StartCityName   string `json:"start_city_name,omitempty"`
    

}
*/

// BusCityChangeDto 
type BusCityChangeDto struct {

    // 是否开通标示，0：已开通 1：未开通
    TypeNo   int64 `json:"type_no,omitempty"`

    // 城市拼音
    CityFullPinyin   string `json:"city_full_pinyin,omitempty"`

    // 是否可售标示，0：暂停售票 1：可售
    Status   int64 `json:"status,omitempty"`

    // 城市名
    StartCityName   string `json:"start_city_name,omitempty"`

}
