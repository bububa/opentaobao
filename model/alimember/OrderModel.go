package alimember

// OrderModel 
type OrderModel struct {
    // 单据类型
    OrderType   string `json:"order_type,omitempty" xml:"order_type,omitempty"`
    // 单据号
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // 该单据对应付费会员开始时间
    OrderIdentityEndTime   string `json:"order_identity_end_time,omitempty" xml:"order_identity_end_time,omitempty"`
    // 该单据对应付费会员结束时间
    OrderIdentityStartTime   string `json:"order_identity_start_time,omitempty" xml:"order_identity_start_time,omitempty"`
}
