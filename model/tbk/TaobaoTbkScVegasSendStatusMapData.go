package tbk

import (
	"sync"
)

// TaobaoTbkScVegasSendStatusMapData 结构体
type TaobaoTbkScVegasSendStatusMapData struct {
	// 若该用户当前无待核销的红包，则返回1，若当前有待核销的红包，则返回0
	IsNewUser string `json:"is_new_user,omitempty" xml:"is_new_user,omitempty"`
}

var poolTaobaoTbkScVegasSendStatusMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScVegasSendStatusMapData)
	},
}

// GetTaobaoTbkScVegasSendStatusMapData() 从对象池中获取TaobaoTbkScVegasSendStatusMapData
func GetTaobaoTbkScVegasSendStatusMapData() *TaobaoTbkScVegasSendStatusMapData {
	return poolTaobaoTbkScVegasSendStatusMapData.Get().(*TaobaoTbkScVegasSendStatusMapData)
}

// ReleaseTaobaoTbkScVegasSendStatusMapData 释放TaobaoTbkScVegasSendStatusMapData
func ReleaseTaobaoTbkScVegasSendStatusMapData(v *TaobaoTbkScVegasSendStatusMapData) {
	v.IsNewUser = ""
	poolTaobaoTbkScVegasSendStatusMapData.Put(v)
}
