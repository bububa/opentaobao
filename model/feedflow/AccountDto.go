package feedflow

import (
	"sync"
)

// AccountDto 结构体
type AccountDto struct {
	// 账户余额，单位：元
	Balance string `json:"balance,omitempty" xml:"balance,omitempty"`
	// 现金余额，单位：元
	CashBalance string `json:"cash_balance,omitempty" xml:"cash_balance,omitempty"`
	// 可用余额，单位：元
	AvailableBalance string `json:"available_balance,omitempty" xml:"available_balance,omitempty"`
	// 红包，单位：元
	RedPacket string `json:"red_packet,omitempty" xml:"red_packet,omitempty"`
}

var poolAccountDto = sync.Pool{
	New: func() any {
		return new(AccountDto)
	},
}

// GetAccountDto() 从对象池中获取AccountDto
func GetAccountDto() *AccountDto {
	return poolAccountDto.Get().(*AccountDto)
}

// ReleaseAccountDto 释放AccountDto
func ReleaseAccountDto(v *AccountDto) {
	v.Balance = ""
	v.CashBalance = ""
	v.AvailableBalance = ""
	v.RedPacket = ""
	poolAccountDto.Put(v)
}
