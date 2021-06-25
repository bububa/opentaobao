package alsc

// AccountOpenInfo 
type AccountOpenInfo struct {

    // 账户ID
    AccountId   string `json:"account_id,omitempty"`

    // 账户类型（暂时不用）
    AccountType   string `json:"account_type,omitempty"`

    // 储值余额
    Balance   int64 `json:"balance,omitempty"`

    // 创建人
    CreateBy   string `json:"create_by,omitempty"`

    // 币种
    Currency   string `json:"currency,omitempty"`

    // 逻辑删除标志
    Deleted   bool `json:"deleted,omitempty"`

    // 增储
    Gift   int64 `json:"gift,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 预储
    Pre   int64 `json:"pre,omitempty"`

    // 实储
    Real   int64 `json:"real,omitempty"`

    // 账户状态（暂时不用）
    Status   string `json:"status,omitempty"`

    // 累计增储
    TotalGift   int64 `json:"total_gift,omitempty"`

    // 累计预储
    TotalPre   int64 `json:"total_pre,omitempty"`

    // 累计实储
    TotalReal   int64 `json:"total_real,omitempty"`

    // 修改人
    UpdateBy   string `json:"update_by,omitempty"`

    // 运营计划ID
    OptPlanId   string `json:"opt_plan_id,omitempty"`

    // 扩展信息
    ExtInfo   *Extinfo `json:"ext_info,omitempty"`

}
