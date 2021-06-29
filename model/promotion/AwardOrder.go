package promotion

// AwardOrder 
type AwardOrder struct {
    // 主订单ID
    BizMainOrderId   int64 `json:"biz_main_order_id,omitempty" xml:"biz_main_order_id,omitempty"`
    // 是否中奖
    IsAward   bool `json:"is_award,omitempty" xml:"is_award,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 排序码
    SortCode   string `json:"sort_code,omitempty" xml:"sort_code,omitempty"`
    // 排序序号
    SortNum   int64 `json:"sort_num,omitempty" xml:"sort_num,omitempty"`
    // 子订单ID
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
