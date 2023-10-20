package waybill

import (
	"sync"
)

// WaybillApplySubscriptionInfo 结构体
type WaybillApplySubscriptionInfo struct {
	// CP网点信息及对应的商家的发货信息
	BranchAccountCols []WaybillBranchAccount `json:"branch_account_cols,omitempty" xml:"branch_account_cols>waybill_branch_account,omitempty"`
	// 物流服务商ID
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 物流服务商业务类型 1：直营 2：加盟 3：落地配 4：直营带网点
	CpType int64 `json:"cp_type,omitempty" xml:"cp_type,omitempty"`
}

var poolWaybillApplySubscriptionInfo = sync.Pool{
	New: func() any {
		return new(WaybillApplySubscriptionInfo)
	},
}

// GetWaybillApplySubscriptionInfo() 从对象池中获取WaybillApplySubscriptionInfo
func GetWaybillApplySubscriptionInfo() *WaybillApplySubscriptionInfo {
	return poolWaybillApplySubscriptionInfo.Get().(*WaybillApplySubscriptionInfo)
}

// ReleaseWaybillApplySubscriptionInfo 释放WaybillApplySubscriptionInfo
func ReleaseWaybillApplySubscriptionInfo(v *WaybillApplySubscriptionInfo) {
	v.BranchAccountCols = v.BranchAccountCols[:0]
	v.CpCode = ""
	v.CpType = 0
	poolWaybillApplySubscriptionInfo.Put(v)
}
