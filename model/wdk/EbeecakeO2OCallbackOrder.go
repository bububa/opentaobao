package wdk

// EbeecakeO2OCallbackOrder 结构体
type EbeecakeO2OCallbackOrder struct {
	// 作业单元列表
	CallbackUnits []EbeecakeO2OCallbackUnit `json:"callback_units,omitempty" xml:"callback_units>ebeecake_o2o_callback_unit,omitempty"`
	// 作业单号
	WorkOrderId string `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
	// 作业单类型： BATCH(&#34;批次&#34;), ORDER(&#34;物流单&#34;)
	WorkOrderType string `json:"work_order_type,omitempty" xml:"work_order_type,omitempty"`
	// 作业状态变更类型：SHIP(&#34;揽收&#34;),SIGN(&#34;妥投&#34;),SIGN_ERROR(&#34;配送异常&#34;),REFUSE(&#34;拒收&#34;)
	StatusChangeType string `json:"status_change_type,omitempty" xml:"status_change_type,omitempty"`
	// 作业状态变更时间
	StatusChangeTime string `json:"status_change_time,omitempty" xml:"status_change_time,omitempty"`
	// 配送员
	Postman *Postman `json:"postman,omitempty" xml:"postman,omitempty"`
}
