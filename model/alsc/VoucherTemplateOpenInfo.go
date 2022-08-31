package alsc

// VoucherTemplateOpenInfo 结构体
type VoucherTemplateOpenInfo struct {
	// 1
	ItemSelectedOpenInfoList []ItemSelectedOpenInfo `json:"item_selected_open_info_list,omitempty" xml:"item_selected_open_info_list>item_selected_open_info,omitempty"`
	// 1
	ShopSelectedOpenInfoList []ShopSelectedOpenInfo `json:"shop_selected_open_info_list,omitempty" xml:"shop_selected_open_info_list>shop_selected_open_info,omitempty"`
	// {\&#34;type\&#34;:0,\&#34;settings\&#34;:[{\&#34;dayStartTime\&#34;:\&#34;00:00\&#34;,\&#34;dayEndTime\&#34;:\&#34;23:59\&#34;,\&#34;week\&#34;:[]}]} type:0不限制，1限制 dayStartTime:开始时间 dayEndTime:结束时间 week:1,2,3,4,5,6,7
	AvailableTime string `json:"available_time,omitempty" xml:"available_time,omitempty"`
	// 使用说明
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 绝对日期中的结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 适用商品范围 值：ALL，PART_AVAILABLE，PART_UNAVAILABLE * 说明：全部商品可用，部分商品可用，部分商品不可用
	ItemCoverage string `json:"item_coverage,omitempty" xml:"item_coverage,omitempty"`
	// 券模板名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 绝对日期中的开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// // 未使用     UNUSED,     // 使用中     USED,     // 无库存     NO_INVENTORY,     // 已失效     INVALID,     ;
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 券模板类型 * 值：GIFT、DISCOUNT、CASH      * 说明：礼品券、礼品券、礼品券
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 使用场景DINE_IN 堂食,TAKE_OUT外卖
	UseCondition string `json:"use_condition,omitempty" xml:"use_condition,omitempty"`
	// 券有效期类型 值： FIXED、RELATIVE      * 说明： 绝对时间，相对于用户领券时间
	ValidDateType string `json:"valid_date_type,omitempty" xml:"valid_date_type,omitempty"`
	// 模版ID
	VoucherTemplateId string `json:"voucher_template_id,omitempty" xml:"voucher_template_id,omitempty"`
	// 面额
	Denomination string `json:"denomination,omitempty" xml:"denomination,omitempty"`
	// 更新人
	UpdateBy string `json:"update_by,omitempty" xml:"update_by,omitempty"`
	// 创建人
	CreateBy string `json:"create_by,omitempty" xml:"create_by,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 库存数量
	Inventory int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// 最低消费
	MinCharge int64 `json:"min_charge,omitempty" xml:"min_charge,omitempty"`
	// 每人限领
	UserLimit int64 `json:"user_limit,omitempty" xml:"user_limit,omitempty"`
	// 相对日期中的有效天数
	ValidDayCount int64 `json:"valid_day_count,omitempty" xml:"valid_day_count,omitempty"`
	// 相对日期中的当天是否有效
	TodayAvailable bool `json:"today_available,omitempty" xml:"today_available,omitempty"`
	// 是否删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}
