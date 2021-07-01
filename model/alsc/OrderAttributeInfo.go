package alsc

// OrderAttributeInfo 结构体
type OrderAttributeInfo struct {
	// 属性所属组:ORDER 订单类，SUBORDER 子订单类
	AttrGroup string `json:"attr_group,omitempty" xml:"attr_group,omitempty"`
	// 套餐明细
	ComboInfoList []ComboInfo `json:"combo_info_list,omitempty" xml:"combo_info_list>combo_info,omitempty"`
	// 单品做法明细
	CookingMethodsInfoList []AttachInfo `json:"cooking_methods_info_list,omitempty" xml:"cooking_methods_info_list>attach_info,omitempty"`
	// 单品配料明细
	IngredientsInfoList []AttachInfo `json:"ingredients_info_list,omitempty" xml:"ingredients_info_list>attach_info,omitempty"`
	// 业务ID,取值与属性所属组对应
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 业务方订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 服务费明细
	ServiceFeeInfoList []ServiceFeeInfo `json:"service_fee_info_list,omitempty" xml:"service_fee_info_list>service_fee_info,omitempty"`
	// 其他附加费，如餐盒等等
	OtherInfoList []AttachInfo `json:"other_info_list,omitempty" xml:"other_info_list>attach_info,omitempty"`
}
