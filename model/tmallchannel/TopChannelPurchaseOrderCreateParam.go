package tmallchannel

// TopChannelPurchaseOrderCreateParam 
type TopChannelPurchaseOrderCreateParam struct {
    // 请求编码
    RequestNo   string `json:"request_no,omitempty" xml:"request_no,omitempty"`
    // 采购明细
    Items   []TopChannelSubPurchaseOrderCreateParam `json:"items,omitempty" xml:"items>top_channel_sub_purchase_order_create_param,omitempty"`
    // 内部编码
    InternalCode   string `json:"internal_code,omitempty" xml:"internal_code,omitempty"`
    // 渠道编码，11-线下网批
    Channel   int64 `json:"channel,omitempty" xml:"channel,omitempty"`
    // 分销商昵称
    DownUserNick   string `json:"down_user_nick,omitempty" xml:"down_user_nick,omitempty"`
    // 是否自动审批
    AutoAudit   bool `json:"auto_audit,omitempty" xml:"auto_audit,omitempty"`
    // 分销商淘宝数字ID，如为空，down_user_nick必须输入
    DownAccountId   int64 `json:"down_account_id,omitempty" xml:"down_account_id,omitempty"`
    // 分销商渠道角色，默认分销终端商
    DownRoleType   int64 `json:"down_role_type,omitempty" xml:"down_role_type,omitempty"`
    // 分销商用户类型，默认淘宝用户
    DownAccountType   int64 `json:"down_account_type,omitempty" xml:"down_account_type,omitempty"`
}
