package idle

import (
	"sync"
)

// TenderPrePayCmd 结构体
type TenderPrePayCmd struct {
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 预付款流水号
	PrePayNo string `json:"pre_pay_no,omitempty" xml:"pre_pay_no,omitempty"`
	// 预付款执行动作
	PrePayAction int64 `json:"pre_pay_action,omitempty" xml:"pre_pay_action,omitempty"`
}

var poolTenderPrePayCmd = sync.Pool{
	New: func() any {
		return new(TenderPrePayCmd)
	},
}

// GetTenderPrePayCmd() 从对象池中获取TenderPrePayCmd
func GetTenderPrePayCmd() *TenderPrePayCmd {
	return poolTenderPrePayCmd.Get().(*TenderPrePayCmd)
}

// ReleaseTenderPrePayCmd 释放TenderPrePayCmd
func ReleaseTenderPrePayCmd(v *TenderPrePayCmd) {
	v.OrderId = ""
	v.PrePayNo = ""
	v.PrePayAction = 0
	poolTenderPrePayCmd.Put(v)
}
