package tbk

import (
	"sync"
)

// TaobaoTbkScPublisherInfoSaveData 结构体
type TaobaoTbkScPublisherInfoSaveData struct {
	// 渠道昵称
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 如果重复绑定会提示：”重复绑定渠道“或”重复绑定粉丝“
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 渠道关系ID
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 会员运营ID
	SpecialId int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
}

var poolTaobaoTbkScPublisherInfoSaveData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScPublisherInfoSaveData)
	},
}

// GetTaobaoTbkScPublisherInfoSaveData() 从对象池中获取TaobaoTbkScPublisherInfoSaveData
func GetTaobaoTbkScPublisherInfoSaveData() *TaobaoTbkScPublisherInfoSaveData {
	return poolTaobaoTbkScPublisherInfoSaveData.Get().(*TaobaoTbkScPublisherInfoSaveData)
}

// ReleaseTaobaoTbkScPublisherInfoSaveData 释放TaobaoTbkScPublisherInfoSaveData
func ReleaseTaobaoTbkScPublisherInfoSaveData(v *TaobaoTbkScPublisherInfoSaveData) {
	v.AccountName = ""
	v.Desc = ""
	v.RelationId = 0
	v.SpecialId = 0
	poolTaobaoTbkScPublisherInfoSaveData.Put(v)
}
