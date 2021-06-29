package subuser

// SubUserInfo 
type SubUserInfo struct {
    // 子账号用户名
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    // 子账号所属的主账号的唯一标识
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    // 主账号昵称
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 子账号当前状态 1正常 -1删除  2冻结
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 是否参与分流 1不参与 2参与
    IsOnline   int64 `json:"is_online,omitempty" xml:"is_online,omitempty"`
    // 子账号Id
    SubId   int64 `json:"sub_id,omitempty" xml:"sub_id,omitempty"`
}
