package hotel

// HotelListInfo 结构体
type HotelListInfo struct {
	// 酒店标签
	Labels []HotelLabelVo `json:"labels,omitempty" xml:"labels>hotel_label_vo,omitempty"`
	// 活动
	PromotionDescArrs []string `json:"promotion_desc_arrs,omitempty" xml:"promotion_desc_arrs>string,omitempty"`
	// service
	Services []Option `json:"services,omitempty" xml:"services>option,omitempty"`
	// 商圈
	Areas []Area `json:"areas,omitempty" xml:"areas>area,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 评论分数
	CommentScore string `json:"comment_score,omitempty" xml:"comment_score,omitempty"`
	// 评论分数的描述文案
	CommentScoreDesc string `json:"comment_score_desc,omitempty" xml:"comment_score_desc,omitempty"`
	// 距离文案
	DistanceDesc string `json:"distance_desc,omitempty" xml:"distance_desc,omitempty"`
	// 英文名
	EnglishName string `json:"english_name,omitempty" xml:"english_name,omitempty"`
	// 酒店fax
	Fax string `json:"fax,omitempty" xml:"fax,omitempty"`
	// 头图地址
	HeaderPicUrl string `json:"header_pic_url,omitempty" xml:"header_pic_url,omitempty"`
	// 酒店经度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 酒店纬度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 酒店名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 酒店星级类型
	StarType string `json:"star_type,omitempty" xml:"star_type,omitempty"`
	// 酒店电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 品牌Id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 品牌名字
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 品牌英文名字
	BrandNameEnglish string `json:"brand_name_english,omitempty" xml:"brand_name_english,omitempty"`
	// inRightMapHotelTitle
	InRightMapHotelTitle string `json:"in_right_map_hotel_title,omitempty" xml:"in_right_map_hotel_title,omitempty"`
	// laterPayIcon
	LaterPayIcon string `json:"later_pay_icon,omitempty" xml:"later_pay_icon,omitempty"`
	// inventoryDesc
	InventoryDesc string `json:"inventory_desc,omitempty" xml:"inventory_desc,omitempty"`
	// 酒店特色id
	FeatureHotelType string `json:"feature_hotel_type,omitempty" xml:"feature_hotel_type,omitempty"`
	// 酒店特色名称
	FeatureHotelTypeName string `json:"feature_hotel_type_name,omitempty" xml:"feature_hotel_type_name,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 评论条数
	CommentCount int64 `json:"comment_count,omitempty" xml:"comment_count,omitempty"`
	// 评论分数
	CommentSource int64 `json:"comment_source,omitempty" xml:"comment_source,omitempty"`
	// 酒店报价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// searchPOI
	SearchPOI *SearchPoi `json:"search_p_o_i,omitempty" xml:"search_p_o_i,omitempty"`
	// 是否有信用住商品
	HasAliCreditItem bool `json:"has_ali_credit_item,omitempty" xml:"has_ali_credit_item,omitempty"`
	// 是否有全景图
	HasFullScenePicture bool `json:"has_full_scene_picture,omitempty" xml:"has_full_scene_picture,omitempty"`
	// 是否售完
	SoldOut bool `json:"sold_out,omitempty" xml:"sold_out,omitempty"`
}
