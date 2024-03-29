package alsc

import (
	"sync"
)

// RechargeAccountOpenInfo 结构体
type RechargeAccountOpenInfo struct {
	// 账户id
	AccountId string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// crm卡实例id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 账户所属顾客id，如果是会员卡，不为空
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 关联的运营方案id
	OptPlanId string `json:"opt_plan_id,omitempty" xml:"opt_plan_id,omitempty"`
	// 账户赠储余额
	GiftValue int64 `json:"gift_value,omitempty" xml:"gift_value,omitempty"`
	// 累计账户赠储余额
	GiftValueTotal int64 `json:"gift_value_total,omitempty" xml:"gift_value_total,omitempty"`
	// 账户预储余额
	PreValue int64 `json:"pre_value,omitempty" xml:"pre_value,omitempty"`
	// 累计账户预储余额
	PreValueTotal int64 `json:"pre_value_total,omitempty" xml:"pre_value_total,omitempty"`
	// 账户实储余额
	RealValue int64 `json:"real_value,omitempty" xml:"real_value,omitempty"`
	// 累计账户实储余额
	RealValueTotal int64 `json:"real_value_total,omitempty" xml:"real_value_total,omitempty"`
	// 可用余额
	UsableValue int64 `json:"usable_value,omitempty" xml:"usable_value,omitempty"`
	// 是否删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolRechargeAccountOpenInfo = sync.Pool{
	New: func() any {
		return new(RechargeAccountOpenInfo)
	},
}

// GetRechargeAccountOpenInfo() 从对象池中获取RechargeAccountOpenInfo
func GetRechargeAccountOpenInfo() *RechargeAccountOpenInfo {
	return poolRechargeAccountOpenInfo.Get().(*RechargeAccountOpenInfo)
}

// ReleaseRechargeAccountOpenInfo 释放RechargeAccountOpenInfo
func ReleaseRechargeAccountOpenInfo(v *RechargeAccountOpenInfo) {
	v.AccountId = ""
	v.CardId = ""
	v.CustomerId = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.OptPlanId = ""
	v.GiftValue = 0
	v.GiftValueTotal = 0
	v.PreValue = 0
	v.PreValueTotal = 0
	v.RealValue = 0
	v.RealValueTotal = 0
	v.UsableValue = 0
	v.Deleted = false
	poolRechargeAccountOpenInfo.Put(v)
}
