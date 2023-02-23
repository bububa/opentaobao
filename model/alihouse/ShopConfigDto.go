package alihouse

// ShopConfigDto 结构体
type ShopConfigDto struct {
	// 配置详细信息
	ShopConfigDetails []ShopConfigDetailDto `json:"shop_config_details,omitempty" xml:"shop_config_details>shop_config_detail_dto,omitempty"`
	// 外部店铺配置ID
	OuterShopConfigId string `json:"outer_shop_config_id,omitempty" xml:"outer_shop_config_id,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 是否为草稿
	IsDraft int64 `json:"is_draft,omitempty" xml:"is_draft,omitempty"`
	// 模块类型
	ModuleType int64 `json:"module_type,omitempty" xml:"module_type,omitempty"`
	// 是否优先展示
	IsPreShow int64 `json:"is_pre_show,omitempty" xml:"is_pre_show,omitempty"`
	// 页面类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 是否展示
	IsShow int64 `json:"is_show,omitempty" xml:"is_show,omitempty"`
	// 是否为测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
