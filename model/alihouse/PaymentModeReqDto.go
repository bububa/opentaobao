package alihouse

// PaymentModeReqDto 结构体
type PaymentModeReqDto struct {
	// 外部楼盘id
	OuterProjectId string `json:"outer_project_id,omitempty" xml:"outer_project_id,omitempty"`
	// 付款方式名称
	PaymentName string `json:"payment_name,omitempty" xml:"payment_name,omitempty"`
	// 创建人姓名
	CreatUserName string `json:"creat_user_name,omitempty" xml:"creat_user_name,omitempty"`
	// 外部付款方式id
	OuterPaymentId string `json:"outer_payment_id,omitempty" xml:"outer_payment_id,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 付款方式类型
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 首套最低比例
	BuyOneRadio int64 `json:"buy_one_radio,omitempty" xml:"buy_one_radio,omitempty"`
	// 多套最低比例
	BuyMultiRadio int64 `json:"buy_multi_radio,omitempty" xml:"buy_multi_radio,omitempty"`
	// 是否测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 是否删除
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}
