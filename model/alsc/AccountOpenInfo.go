package alsc

import (
	"sync"
)

// AccountOpenInfo 结构体
type AccountOpenInfo struct {
	// 账户ID
	AccountId string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 账户类型（暂时不用）
	AccountType string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 创建人
	CreateBy string `json:"create_by,omitempty" xml:"create_by,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 账户状态（暂时不用）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 修改人
	UpdateBy string `json:"update_by,omitempty" xml:"update_by,omitempty"`
	// 运营计划ID
	OptPlanId string `json:"opt_plan_id,omitempty" xml:"opt_plan_id,omitempty"`
	// 储值余额
	Balance int64 `json:"balance,omitempty" xml:"balance,omitempty"`
	// 增储
	Gift int64 `json:"gift,omitempty" xml:"gift,omitempty"`
	// 预储
	Pre int64 `json:"pre,omitempty" xml:"pre,omitempty"`
	// 实储
	Real int64 `json:"real,omitempty" xml:"real,omitempty"`
	// 累计增储
	TotalGift int64 `json:"total_gift,omitempty" xml:"total_gift,omitempty"`
	// 累计预储
	TotalPre int64 `json:"total_pre,omitempty" xml:"total_pre,omitempty"`
	// 累计实储
	TotalReal int64 `json:"total_real,omitempty" xml:"total_real,omitempty"`
	// 扩展信息
	ExtInfo *Extinfo `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 逻辑删除标志
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolAccountOpenInfo = sync.Pool{
	New: func() any {
		return new(AccountOpenInfo)
	},
}

// GetAccountOpenInfo() 从对象池中获取AccountOpenInfo
func GetAccountOpenInfo() *AccountOpenInfo {
	return poolAccountOpenInfo.Get().(*AccountOpenInfo)
}

// ReleaseAccountOpenInfo 释放AccountOpenInfo
func ReleaseAccountOpenInfo(v *AccountOpenInfo) {
	v.AccountId = ""
	v.AccountType = ""
	v.CreateBy = ""
	v.Currency = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Status = ""
	v.UpdateBy = ""
	v.OptPlanId = ""
	v.Balance = 0
	v.Gift = 0
	v.Pre = 0
	v.Real = 0
	v.TotalGift = 0
	v.TotalPre = 0
	v.TotalReal = 0
	v.ExtInfo = nil
	v.Deleted = false
	poolAccountOpenInfo.Put(v)
}
