package product

// ModularDescInfo 结构体
type ModularDescInfo struct {
	// 运营定义的该商品所属类目的类目级别模块信息列表，列表顺序即为模块顺序。
	ItmDscModules []ItemDescModule `json:"itm_dsc_modules,omitempty" xml:"itm_dsc_modules>item_desc_module,omitempty"`
	// 用户自定义模块数量最大值限制。类目级别模块+用户级别模块必须小于<20
	UsrDefMaxNum int64 `json:"usr_def_max_num,omitempty" xml:"usr_def_max_num,omitempty"`
	// 旧描述与新描述共存时间（切换时间）。
	DeadLine string `json:"dead_line,omitempty" xml:"dead_line,omitempty"`
	// 默认值为false，如果此字段为true，则卖家上传的模块化的描述信息可以自定义排序。
	UserOrder bool `json:"user_order,omitempty" xml:"user_order,omitempty"`
}
