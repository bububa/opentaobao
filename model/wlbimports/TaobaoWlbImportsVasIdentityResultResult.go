package wlbimports

import (
	"sync"
)

// TaobaoWlbImportsVasIdentityResultResult 结构体
type TaobaoWlbImportsVasIdentityResultResult struct {
	// 鉴定结果项
	VasResults []IdentityItemDto `json:"vas_results,omitempty" xml:"vas_results>identity_item_dto,omitempty"`
	// 物流订单号
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
}

var poolTaobaoWlbImportsVasIdentityResultResult = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportsVasIdentityResultResult)
	},
}

// GetTaobaoWlbImportsVasIdentityResultResult() 从对象池中获取TaobaoWlbImportsVasIdentityResultResult
func GetTaobaoWlbImportsVasIdentityResultResult() *TaobaoWlbImportsVasIdentityResultResult {
	return poolTaobaoWlbImportsVasIdentityResultResult.Get().(*TaobaoWlbImportsVasIdentityResultResult)
}

// ReleaseTaobaoWlbImportsVasIdentityResultResult 释放TaobaoWlbImportsVasIdentityResultResult
func ReleaseTaobaoWlbImportsVasIdentityResultResult(v *TaobaoWlbImportsVasIdentityResultResult) {
	v.VasResults = v.VasResults[:0]
	v.LgOrderCode = ""
	poolTaobaoWlbImportsVasIdentityResultResult.Put(v)
}
