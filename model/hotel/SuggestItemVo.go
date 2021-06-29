package hotel

// SuggestItemVo 
type SuggestItemVo struct {
    // 城市code
    CityCode   int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
    // 城市名字
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    // 联想词
    DisplayName   string `json:"display_name,omitempty" xml:"display_name,omitempty"`
    // 联想词类型
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 英文的联想词
    DisplayNameEnglish   string `json:"display_name_english,omitempty" xml:"display_name_english,omitempty"`
    // 区域，国内0，国外1
    Region   int64 `json:"region,omitempty" xml:"region,omitempty"`
}
