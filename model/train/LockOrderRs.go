package train

import (
	"sync"
)

// LockOrderRs 结构体
type LockOrderRs struct {
	// 票信息
	TicketInfos []TicketInfoDto `json:"ticket_infos,omitempty" xml:"ticket_infos>ticket_info_dto,omitempty"`
	// 锁单最晚有效时间
	LockLastTime string `json:"lock_last_time,omitempty" xml:"lock_last_time,omitempty"`
	// 乘车人余留联系方式
	ContactMobileNo string `json:"contact_mobile_no,omitempty" xml:"contact_mobile_no,omitempty"`
}

var poolLockOrderRs = sync.Pool{
	New: func() any {
		return new(LockOrderRs)
	},
}

// GetLockOrderRs() 从对象池中获取LockOrderRs
func GetLockOrderRs() *LockOrderRs {
	return poolLockOrderRs.Get().(*LockOrderRs)
}

// ReleaseLockOrderRs 释放LockOrderRs
func ReleaseLockOrderRs(v *LockOrderRs) {
	v.TicketInfos = v.TicketInfos[:0]
	v.LockLastTime = ""
	v.ContactMobileNo = ""
	poolLockOrderRs.Put(v)
}
