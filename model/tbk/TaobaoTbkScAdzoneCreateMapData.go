package tbk

import (
	"sync"
)

// TaobaoTbkScAdzoneCreateMapData 结构体
type TaobaoTbkScAdzoneCreateMapData struct {
	// 完整的pid
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}

var poolTaobaoTbkScAdzoneCreateMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScAdzoneCreateMapData)
	},
}

// GetTaobaoTbkScAdzoneCreateMapData() 从对象池中获取TaobaoTbkScAdzoneCreateMapData
func GetTaobaoTbkScAdzoneCreateMapData() *TaobaoTbkScAdzoneCreateMapData {
	return poolTaobaoTbkScAdzoneCreateMapData.Get().(*TaobaoTbkScAdzoneCreateMapData)
}

// ReleaseTaobaoTbkScAdzoneCreateMapData 释放TaobaoTbkScAdzoneCreateMapData
func ReleaseTaobaoTbkScAdzoneCreateMapData(v *TaobaoTbkScAdzoneCreateMapData) {
	v.Model = ""
	poolTaobaoTbkScAdzoneCreateMapData.Put(v)
}
