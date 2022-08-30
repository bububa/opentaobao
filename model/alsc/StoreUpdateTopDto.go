package alsc

// StoreUpdateTopDto 结构体
type StoreUpdateTopDto struct {
	// 门店标
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 通用属性
	Attributes []AttributeValueTopDto `json:"attributes,omitempty" xml:"attributes>attribute_value_top_dto,omitempty"`
	// 类目属性
	CategoryPropertyValues []PropertyValueTopDto `json:"category_property_values,omitempty" xml:"category_property_values>property_value_top_dto,omitempty"`
	// 业务属性
	BizAttributes []AttributeValueTopDto `json:"biz_attributes,omitempty" xml:"biz_attributes>attribute_value_top_dto,omitempty"`
	// 星期
	Week []string `json:"week,omitempty" xml:"week>string,omitempty"`
	// 备注
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 门店状态，枚举值。NORMAL：正常。CLOSE：关店。HOLD: 暂停营业
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 门店结束营业时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 门店开始营业时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 门店主名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 门店外部编码
	OuterCode string `json:"outer_code,omitempty" xml:"outer_code,omitempty"`
	// 门店类型，枚举值。NORMAL：普通门店。暂时统一使用这个值
	StoreType string `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// 分店名称
	SubName string `json:"sub_name,omitempty" xml:"sub_name,omitempty"`
	// 标准类目ID
	StandardCategoryId string `json:"standard_category_id,omitempty" xml:"standard_category_id,omitempty"`
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 门店logo
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 门店头图
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 门店信息
	StoreKeeper *StoreKeeperDto `json:"store_keeper,omitempty" xml:"store_keeper,omitempty"`
	// 门店地址
	StoreAdress *StoreAdressDto `json:"store_adress,omitempty" xml:"store_adress,omitempty"`
	// 门店主类目
	MainCategory int64 `json:"main_category,omitempty" xml:"main_category,omitempty"`
	// 店铺id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 门店类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 审核状态
	AuthenStatus int64 `json:"authen_status,omitempty" xml:"authen_status,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 24小时营业
	AllDay bool `json:"all_day,omitempty" xml:"all_day,omitempty"`
	// 是否v3
	IsV3 bool `json:"is_v3,omitempty" xml:"is_v3,omitempty"`
}
