package tbk

import (
	"sync"
)

// TaobaoTbkScPunishOrderGetResult 结构体
type TaobaoTbkScPunishOrderGetResult struct {
	// 淘宝联盟unionid（该字段不再支持）
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// 处罚状态。0 冻结，1 解冻
	PunishStatus string `json:"punish_status,omitempty" xml:"punish_status,omitempty"`
	// 处罚类型，目前包括 1:店铺淘宝客 2:订单虚假交易
	ViolationType string `json:"violation_type,omitempty" xml:"violation_type,omitempty"`
	// 淘客订单创建时间
	TkTradeCreateTime string `json:"tk_trade_create_time,omitempty" xml:"tk_trade_create_time,omitempty"`
	// 会员运营id（该字段不再支持）
	SpecialId int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
	// 渠道关系id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 结算月份
	SettleMonth int64 `json:"settle_month,omitempty" xml:"settle_month,omitempty"`
	// 子订单号
	TbTradeId int64 `json:"tb_trade_id,omitempty" xml:"tb_trade_id,omitempty"`
	// 父订单号（该字段不再支持）
	TbTradeParentId int64 `json:"tb_trade_parent_id,omitempty" xml:"tb_trade_parent_id,omitempty"`
	// pid里的adzoneid
	TkAdzoneId int64 `json:"tk_adzone_id,omitempty" xml:"tk_adzone_id,omitempty"`
	// pid里的siteid
	TkSiteId int64 `json:"tk_site_id,omitempty" xml:"tk_site_id,omitempty"`
	// pid里的pubid
	TkPubId int64 `json:"tk_pub_id,omitempty" xml:"tk_pub_id,omitempty"`
}

var poolTaobaoTbkScPunishOrderGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScPunishOrderGetResult)
	},
}

// GetTaobaoTbkScPunishOrderGetResult() 从对象池中获取TaobaoTbkScPunishOrderGetResult
func GetTaobaoTbkScPunishOrderGetResult() *TaobaoTbkScPunishOrderGetResult {
	return poolTaobaoTbkScPunishOrderGetResult.Get().(*TaobaoTbkScPunishOrderGetResult)
}

// ReleaseTaobaoTbkScPunishOrderGetResult 释放TaobaoTbkScPunishOrderGetResult
func ReleaseTaobaoTbkScPunishOrderGetResult(v *TaobaoTbkScPunishOrderGetResult) {
	v.UnionId = ""
	v.PunishStatus = ""
	v.ViolationType = ""
	v.TkTradeCreateTime = ""
	v.SpecialId = 0
	v.RelationId = 0
	v.SettleMonth = 0
	v.TbTradeId = 0
	v.TbTradeParentId = 0
	v.TkAdzoneId = 0
	v.TkSiteId = 0
	v.TkPubId = 0
	poolTaobaoTbkScPunishOrderGetResult.Put(v)
}
