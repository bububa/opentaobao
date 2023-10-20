package ascpqcc

import (
	"sync"
)

// CancelSampleRelationData 结构体
type CancelSampleRelationData struct {
	// 商品编号列表
	BizItemIds []string `json:"biz_item_ids,omitempty" xml:"biz_item_ids>string,omitempty"`
	// 商品编号
	BizItemId string `json:"biz_item_id,omitempty" xml:"biz_item_id,omitempty"`
}

var poolCancelSampleRelationData = sync.Pool{
	New: func() any {
		return new(CancelSampleRelationData)
	},
}

// GetCancelSampleRelationData() 从对象池中获取CancelSampleRelationData
func GetCancelSampleRelationData() *CancelSampleRelationData {
	return poolCancelSampleRelationData.Get().(*CancelSampleRelationData)
}

// ReleaseCancelSampleRelationData 释放CancelSampleRelationData
func ReleaseCancelSampleRelationData(v *CancelSampleRelationData) {
	v.BizItemIds = v.BizItemIds[:0]
	v.BizItemId = ""
	poolCancelSampleRelationData.Put(v)
}
