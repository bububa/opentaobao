package btrip

// CityVo 
type CityVo struct {
    // 三字码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 城市名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 城市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 所属省份
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    // 与搜索城市距离，单位千米，只在邻近机场推荐有值
    Distance   int64 `json:"distance,omitempty" xml:"distance,omitempty"`
    // 邻近机场城市，只在邻近机场推荐有值
    TravelName   string `json:"travel_name,omitempty" xml:"travel_name,omitempty"`
}
