package crm

import (
	"sync"
)

// BasicMember 结构体
type BasicMember struct {
	// 会员昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// open_uid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// 交易关闭金额
	CloseTradeAmount string `json:"close_trade_amount,omitempty" xml:"close_trade_amount,omitempty"`
	// 分组的id集合字符串
	GroupIds string `json:"group_ids,omitempty" xml:"group_ids,omitempty"`
	// 显示会员的状态，normal正常，blacklist黑名单
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 交易的金额
	TradeAmount string `json:"trade_amount,omitempty" xml:"trade_amount,omitempty"`
	// 最后交易的日期
	LastTradeTime string `json:"last_trade_time,omitempty" xml:"last_trade_time,omitempty"`
	// 等级名称
	GradeName string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
	// 购买的宝贝件数
	ItemNum int64 `json:"item_num,omitempty" xml:"item_num,omitempty"`
	// 关系来源，1交易成功，2未交易成功 ,3 卖家主动吸纳
	RelationSource int64 `json:"relation_source,omitempty" xml:"relation_source,omitempty"`
	// 会员等级，0未非会员，其余对应等级名称取grade_name
	Grade int64 `json:"grade,omitempty" xml:"grade,omitempty"`
	// 交易关闭的笔数
	CloseTradeCount int64 `json:"close_trade_count,omitempty" xml:"close_trade_count,omitempty"`
	// 交易成功的次数
	TradeCount int64 `json:"trade_count,omitempty" xml:"trade_count,omitempty"`
	// 最后一次交易的订单号.注:该字段从2014.4.23之后不再返回.
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolBasicMember = sync.Pool{
	New: func() any {
		return new(BasicMember)
	},
}

// GetBasicMember() 从对象池中获取BasicMember
func GetBasicMember() *BasicMember {
	return poolBasicMember.Get().(*BasicMember)
}

// ReleaseBasicMember 释放BasicMember
func ReleaseBasicMember(v *BasicMember) {
	v.BuyerNick = ""
	v.Ouid = ""
	v.CloseTradeAmount = ""
	v.GroupIds = ""
	v.Status = ""
	v.TradeAmount = ""
	v.LastTradeTime = ""
	v.GradeName = ""
	v.ItemNum = 0
	v.RelationSource = 0
	v.Grade = 0
	v.CloseTradeCount = 0
	v.TradeCount = 0
	v.BizOrderId = 0
	poolBasicMember.Put(v)
}
