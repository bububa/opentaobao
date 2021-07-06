package store

// FullStoreTopDto 结构体
type FullStoreTopDto struct {
	// 业务TAGS
	Tags []int64 `json:"tags,omitempty" xml:"tags>int64,omitempty"`
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 门店名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 业务外部id
	BizOuterId string `json:"biz_outer_id,omitempty" xml:"biz_outer_id,omitempty"`
	// 分店名
	Subname string `json:"subname,omitempty" xml:"subname,omitempty"`
	// 门店所属商户名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 门店描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 门店LOGO
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 门店外部编码
	OuterCode string `json:"outer_code,omitempty" xml:"outer_code,omitempty"`
	// 门店状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 门店类型
	StoreType string `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// 门店主图，仅存一张，tfs格式
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 父级别门店id
	Belong int64 `json:"belong,omitempty" xml:"belong,omitempty"`
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 门店类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 门店的商家id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 审核状态
	AuthenStatus int64 `json:"authen_status,omitempty" xml:"authen_status,omitempty"`
	// 位置地址信息
	PoiInfo *PoiInfoDto `json:"poi_info,omitempty" xml:"poi_info,omitempty"`
	// 店铺id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 标准类目ID
	StandardCategoryId int64 `json:"standard_category_id,omitempty" xml:"standard_category_id,omitempty"`
}
