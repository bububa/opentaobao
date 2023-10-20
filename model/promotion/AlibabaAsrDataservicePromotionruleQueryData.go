package promotion

import (
	"sync"
)

// AlibabaAsrDataservicePromotionruleQueryData 结构体
type AlibabaAsrDataservicePromotionruleQueryData struct {
	// 兑换详情列表
	DetailList []Detaillist `json:"detail_list,omitempty" xml:"detail_list>detaillist,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// poskey
	PosKey int64 `json:"pos_key,omitempty" xml:"pos_key,omitempty"`
}

var poolAlibabaAsrDataservicePromotionruleQueryData = sync.Pool{
	New: func() any {
		return new(AlibabaAsrDataservicePromotionruleQueryData)
	},
}

// GetAlibabaAsrDataservicePromotionruleQueryData() 从对象池中获取AlibabaAsrDataservicePromotionruleQueryData
func GetAlibabaAsrDataservicePromotionruleQueryData() *AlibabaAsrDataservicePromotionruleQueryData {
	return poolAlibabaAsrDataservicePromotionruleQueryData.Get().(*AlibabaAsrDataservicePromotionruleQueryData)
}

// ReleaseAlibabaAsrDataservicePromotionruleQueryData 释放AlibabaAsrDataservicePromotionruleQueryData
func ReleaseAlibabaAsrDataservicePromotionruleQueryData(v *AlibabaAsrDataservicePromotionruleQueryData) {
	v.DetailList = v.DetailList[:0]
	v.Name = ""
	v.PosKey = 0
	poolAlibabaAsrDataservicePromotionruleQueryData.Put(v)
}
