package tbk

// TaobaoTbkDgPunishOrderGetResult 
type TaobaoTbkDgPunishOrderGetResult struct {

    // 渠道关系id
    RelationId   int64 `json:"relation_id,omitempty"`

    // 结算月份
    SettleMonth   int64 `json:"settle_month,omitempty"`

    // 淘客订单创建时间
    TkTradeCreateTime   string `json:"tk_trade_create_time,omitempty"`

    // 父订单号（该字段不再支持）
    TbTradeParentId   int64 `json:"tb_trade_parent_id,omitempty"`

    // 会员运营id（该字段不再支持）
    SpecialId   int64 `json:"special_id,omitempty"`

    // 淘宝联盟unionid（该字段不再支持）
    UnionId   string `json:"union_id,omitempty"`

    // 处罚状态。0 冻结，1 解冻
    PunishStatus   string `json:"punish_status,omitempty"`

    // 处罚类型，目前包括 1.店铺淘宝客 2.订单虚假交易
    ViolationType   string `json:"violation_type,omitempty"`

    // 子订单号
    TbTradeId   int64 `json:"tb_trade_id,omitempty"`

    // pid里的adzoneid
    TkAdzoneId   int64 `json:"tk_adzone_id,omitempty"`

    // pid里的siteid
    TkSiteId   int64 `json:"tk_site_id,omitempty"`

    // pid里的pubid
    TkPubId   string `json:"tk_pub_id,omitempty"`

}
