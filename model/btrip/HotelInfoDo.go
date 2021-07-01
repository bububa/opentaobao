package btrip

// HotelInfoDo 
type HotelInfoDo struct {
    // 酒店中文名
    HotelName   string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
    // 酒店联系方式
    HotelTel   string `json:"hotel_tel,omitempty" xml:"hotel_tel,omitempty"`
    // 酒店地址
    HotelAddress   string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
    // 所在城市code
    HotelCityCode   string `json:"hotel_city_code,omitempty" xml:"hotel_city_code,omitempty"`
    // 城市名称
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 酒店星级
    StarRating   string `json:"star_rating,omitempty" xml:"star_rating,omitempty"`
}
