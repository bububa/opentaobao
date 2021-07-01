package travel

// FreeTourHotelInfo 结构体
type FreeTourHotelInfo struct {
	// 酒店描述
	HotelDesc string `json:"hotel_desc,omitempty" xml:"hotel_desc,omitempty"`
	// 新发布商品必填，最多三张，图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
	PicUrls string `json:"pic_urls,omitempty" xml:"pic_urls,omitempty"`
	// 酒店所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 必填，酒店房型
	HouseType string `json:"house_type,omitempty" xml:"house_type,omitempty"`
	// 酒店档次（与下面的酒店星级二选一，只能填一个）：经济连锁；舒适；高档；豪华；二星及以下
	HotelLevel string `json:"hotel_level,omitempty" xml:"hotel_level,omitempty"`
	// 酒店星级（与上面的酒店档次二选一，只能填一个）
	HotelStar string `json:"hotel_star,omitempty" xml:"hotel_star,omitempty"`
	// 酒店名称
	CnName string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
}
