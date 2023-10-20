package bill

import (
	"sync"
)

// TopAcctCashJourDto 结构体
type TopAcctCashJourDto struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 备注
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 流水的淘宝支付宝id
	TaobaoAlipayId string `json:"taobao_alipay_id,omitempty" xml:"taobao_alipay_id,omitempty"`
	// 流水的商家支付宝id
	OtherAlipayId string `json:"other_alipay_id,omitempty" xml:"other_alipay_id,omitempty"`
	// 记账时间
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 操作金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 流水类型:101-可用金充值;102-可用金扣除;103-冻结;104-解冻;105-冻结金充值;106-冻结金扣除
	JournalType int64 `json:"journal_type,omitempty" xml:"journal_type,omitempty"`
	// 虚拟账户科目编号
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 虚拟账户流水编号
	Bid int64 `json:"bid,omitempty" xml:"bid,omitempty"`
}

var poolTopAcctCashJourDto = sync.Pool{
	New: func() any {
		return new(TopAcctCashJourDto)
	},
}

// GetTopAcctCashJourDto() 从对象池中获取TopAcctCashJourDto
func GetTopAcctCashJourDto() *TopAcctCashJourDto {
	return poolTopAcctCashJourDto.Get().(*TopAcctCashJourDto)
}

// ReleaseTopAcctCashJourDto 释放TopAcctCashJourDto
func ReleaseTopAcctCashJourDto(v *TopAcctCashJourDto) {
	v.GmtCreate = ""
	v.Description = ""
	v.TaobaoAlipayId = ""
	v.OtherAlipayId = ""
	v.BookTime = ""
	v.Amount = 0
	v.JournalType = 0
	v.AccountId = 0
	v.Bid = 0
	poolTopAcctCashJourDto.Put(v)
}
