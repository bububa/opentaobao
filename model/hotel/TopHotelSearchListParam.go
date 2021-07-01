package hotel

// TopHotelSearchListParam 结构体
type TopHotelSearchListParam struct {
	// 成人数
	AdultNum int64 `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	// 入店时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 儿童年龄
	ChildrenAges []int64 `json:"children_ages,omitempty" xml:"children_ages>int64,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 关键字
	Keywords string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// 酒店获取开始位置
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 排序方向
	OrderDirection int64 `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	// 排序类型
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 每页个数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 距离搜索时的搜索半径
	Radius int64 `json:"radius,omitempty" xml:"radius,omitempty"`
	// 卖家ids
	SellerIds []int64 `json:"seller_ids,omitempty" xml:"seller_ids>int64,omitempty"`
	// 用户的Agent
	UserAgent string `json:"user_agent,omitempty" xml:"user_agent,omitempty"`
	// 用户Id, 0表示未登录用户
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户定位的经度
	UserLatitude string `json:"user_latitude,omitempty" xml:"user_latitude,omitempty"`
	// 用户定位的伟度
	UserLongitude string `json:"user_longitude,omitempty" xml:"user_longitude,omitempty"`
	// 搜索poi
	SearchPoi *SearchPoi `json:"search_poi,omitempty" xml:"search_poi,omitempty"`
	// 品牌code
	BrandCodeList []string `json:"brand_code_list,omitempty" xml:"brand_code_list>string,omitempty"`
	// 用户city
	UserCity int64 `json:"user_city,omitempty" xml:"user_city,omitempty"`
}
