package bill

// Account 
type Account struct {

    // 费用科目编号
    AccountId   int64 `json:"account_id,omitempty"`

    // 费用科目名称
    AccountName   string `json:"account_name,omitempty"`

    // 费用科目编码
    AccountCode   string `json:"account_code,omitempty"`

    // 费用科目类型:1-虚拟账户 2-交易 3-交易佣金 4-服务费 5-天猫积分 6-其他
    AccountType   int64 `json:"account_type,omitempty"`

    // 是否订单相关:0-订单无关 1-订单相关
    RelatedOrder   int64 `json:"related_order,omitempty"`

    // 状态:0-不可用 1-可用
    Status   int64 `json:"status,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

}
