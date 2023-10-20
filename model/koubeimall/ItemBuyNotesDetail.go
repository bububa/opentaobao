package koubeimall

import (
	"sync"
)

// ItemBuyNotesDetail 结构体
type ItemBuyNotesDetail struct {
	// 须知明细
	NotesDetails []string `json:"notes_details,omitempty" xml:"notes_details>string,omitempty"`
	// 须知标题
	NotesDetailTitle string `json:"notes_detail_title,omitempty" xml:"notes_detail_title,omitempty"`
}

var poolItemBuyNotesDetail = sync.Pool{
	New: func() any {
		return new(ItemBuyNotesDetail)
	},
}

// GetItemBuyNotesDetail() 从对象池中获取ItemBuyNotesDetail
func GetItemBuyNotesDetail() *ItemBuyNotesDetail {
	return poolItemBuyNotesDetail.Get().(*ItemBuyNotesDetail)
}

// ReleaseItemBuyNotesDetail 释放ItemBuyNotesDetail
func ReleaseItemBuyNotesDetail(v *ItemBuyNotesDetail) {
	v.NotesDetails = v.NotesDetails[:0]
	v.NotesDetailTitle = ""
	poolItemBuyNotesDetail.Put(v)
}
