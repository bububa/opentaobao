package tbk

import (
	"sync"
)

// TaobaoTbkDgVegasSendStatusMapData 结构体
type TaobaoTbkDgVegasSendStatusMapData struct {
	// 若该用户当前无待核销的红包，则返回1，若当前有待核销的红包，则返回0
	IsNewUser string `json:"is_new_user,omitempty" xml:"is_new_user,omitempty"`
}

var poolTaobaoTbkDgVegasSendStatusMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgVegasSendStatusMapData)
	},
}

// GetTaobaoTbkDgVegasSendStatusMapData() 从对象池中获取TaobaoTbkDgVegasSendStatusMapData
func GetTaobaoTbkDgVegasSendStatusMapData() *TaobaoTbkDgVegasSendStatusMapData {
	return poolTaobaoTbkDgVegasSendStatusMapData.Get().(*TaobaoTbkDgVegasSendStatusMapData)
}

// ReleaseTaobaoTbkDgVegasSendStatusMapData 释放TaobaoTbkDgVegasSendStatusMapData
func ReleaseTaobaoTbkDgVegasSendStatusMapData(v *TaobaoTbkDgVegasSendStatusMapData) {
	v.IsNewUser = ""
	poolTaobaoTbkDgVegasSendStatusMapData.Put(v)
}
