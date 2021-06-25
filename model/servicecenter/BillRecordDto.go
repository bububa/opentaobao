package servicecenter

// BillRecordDto 
type BillRecordDto struct {

    // 记录产生时间
    StartDate   string `json:"start_date,omitempty"`

    // 状态：1成功、2失败
    Status   int64 `json:"status,omitempty"`

    // 备用字段
    Extend1   string `json:"extend1,omitempty"`

    // 备用字段
    Extend10   string `json:"extend10,omitempty"`

    // appkey
    Appkey   string `json:"appkey,omitempty"`

    // 备用字段
    Extend3   string `json:"extend3,omitempty"`

    // 备用字段
    Extend2   string `json:"extend2,omitempty"`

    // 备用字段
    Extend5   string `json:"extend5,omitempty"`

    // 备用字段
    Extend4   string `json:"extend4,omitempty"`

    // 账单分类：1短信
    Type   int64 `json:"type,omitempty"`

    // 目标号码
    TargetNo   string `json:"target_no,omitempty"`

    // 备用字段
    Extend7   string `json:"extend7,omitempty"`

    // 备用字段
    Extend6   string `json:"extend6,omitempty"`

    // 备用字段
    Extend9   string `json:"extend9,omitempty"`

    // 外部确认账单ID
    OutConfirmId   string `json:"out_confirm_id,omitempty"`

    // 备用字段
    Extend8   string `json:"extend8,omitempty"`

    // 金额（单位：分）
    Fee   int64 `json:"fee,omitempty"`

    // 卖家ID
    Nick   string `json:"nick,omitempty"`

    // 服务市场有效订单ID
    OrderId   int64 `json:"order_id,omitempty"`

    // 外部订单ID
    OutOrderId   string `json:"out_order_id,omitempty"`

}
