package ascpchannel

import (
	"sync"
)

// TaobaoInvTurnoverQueryData 结构体
type TaobaoInvTurnoverQueryData struct {
	// 明细
	BussinessInventoryLog *BussinessInventoryLog `json:"bussiness_inventory_log,omitempty" xml:"bussiness_inventory_log,omitempty"`
}

var poolTaobaoInvTurnoverQueryData = sync.Pool{
	New: func() any {
		return new(TaobaoInvTurnoverQueryData)
	},
}

// GetTaobaoInvTurnoverQueryData() 从对象池中获取TaobaoInvTurnoverQueryData
func GetTaobaoInvTurnoverQueryData() *TaobaoInvTurnoverQueryData {
	return poolTaobaoInvTurnoverQueryData.Get().(*TaobaoInvTurnoverQueryData)
}

// ReleaseTaobaoInvTurnoverQueryData 释放TaobaoInvTurnoverQueryData
func ReleaseTaobaoInvTurnoverQueryData(v *TaobaoInvTurnoverQueryData) {
	v.BussinessInventoryLog = nil
	poolTaobaoInvTurnoverQueryData.Put(v)
}
