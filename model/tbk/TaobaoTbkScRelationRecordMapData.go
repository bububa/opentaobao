package tbk

import (
	"sync"
)

// TaobaoTbkScRelationRecordMapData 结构体
type TaobaoTbkScRelationRecordMapData struct {
	// 带授权的备案链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 已废弃，请调用淘口令生成接口
	PasswordSimple string `json:"password_simple,omitempty" xml:"password_simple,omitempty"`
	// 已废弃，请调用淘口令生成接口
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}

var poolTaobaoTbkScRelationRecordMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScRelationRecordMapData)
	},
}

// GetTaobaoTbkScRelationRecordMapData() 从对象池中获取TaobaoTbkScRelationRecordMapData
func GetTaobaoTbkScRelationRecordMapData() *TaobaoTbkScRelationRecordMapData {
	return poolTaobaoTbkScRelationRecordMapData.Get().(*TaobaoTbkScRelationRecordMapData)
}

// ReleaseTaobaoTbkScRelationRecordMapData 释放TaobaoTbkScRelationRecordMapData
func ReleaseTaobaoTbkScRelationRecordMapData(v *TaobaoTbkScRelationRecordMapData) {
	v.Url = ""
	v.PasswordSimple = ""
	v.Model = ""
	poolTaobaoTbkScRelationRecordMapData.Put(v)
}
