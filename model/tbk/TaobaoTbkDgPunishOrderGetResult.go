package tbk

// TaobaotbkdgpunishordergetResult 结构体
type TaobaotbkdgpunishordergetResult struct {
	// 淘宝联盟unionid（该字段不再支持）
	Unionid string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// 处罚状态。0 冻结，1 解冻
	Punishstatus string `json:"punish_status,omitempty" xml:"punish_status,omitempty"`
	// 处罚类型，目前包括 1.店铺淘宝客 2.订单虚假交易
	Violationtype string `json:"violation_type,omitempty" xml:"violation_type,omitempty"`
	// 淘客订单创建时间
	Tktradecreatetime string `json:"tk_trade_create_time,omitempty" xml:"tk_trade_create_time,omitempty"`
	// pid里的pubid
	Tkpubid string `json:"tk_pub_id,omitempty" xml:"tk_pub_id,omitempty"`
	// 会员运营id（该字段不再支持）
	Specialid int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
	// 渠道关系id
	Relationid int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 结算月份
	Settlemonth int64 `json:"settle_month,omitempty" xml:"settle_month,omitempty"`
	// 子订单号
	Tbtradeid int64 `json:"tb_trade_id,omitempty" xml:"tb_trade_id,omitempty"`
	// 父订单号（该字段不再支持）
	Tbtradeparentid int64 `json:"tb_trade_parent_id,omitempty" xml:"tb_trade_parent_id,omitempty"`
	// pid里的adzoneid
	Tkadzoneid int64 `json:"tk_adzone_id,omitempty" xml:"tk_adzone_id,omitempty"`
	// pid里的siteid
	Tksiteid int64 `json:"tk_site_id,omitempty" xml:"tk_site_id,omitempty"`
}
