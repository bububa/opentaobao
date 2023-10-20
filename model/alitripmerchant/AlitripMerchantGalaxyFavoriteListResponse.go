package alitripmerchant

// AlitripmerchantgalaxyfavoritelistResponse 结构体
type AlitripmerchantgalaxyfavoritelistResponse struct {
	// 收藏列表
	Contents []FavoriteHotelList `json:"contents,omitempty" xml:"contents>favorite_hotel_list,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
