package train

// AccountParam 
type AccountParam struct {

    // 代理商ID
    AgentId   int64 `json:"agent_id,omitempty"`

    // 主订单号
    OrderId   int64 `json:"order_id,omitempty"`

    // 12306账户名
    AccountName   string `json:"account_name,omitempty"`

    // 12306账户密码
    AccountPassword   string `json:"account_password,omitempty"`

}
