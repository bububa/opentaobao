package koubeimall

import (
	"sync"
)

// ItemBuyNotes 结构体
type ItemBuyNotes struct {
	// 列表
	BuyNotesDetailList []ItemBuyNotesDetail `json:"buy_notes_detail_list,omitempty" xml:"buy_notes_detail_list>item_buy_notes_detail,omitempty"`
	// 购买须知标题
	NotesTitle string `json:"notes_title,omitempty" xml:"notes_title,omitempty"`
}

var poolItemBuyNotes = sync.Pool{
	New: func() any {
		return new(ItemBuyNotes)
	},
}

// GetItemBuyNotes() 从对象池中获取ItemBuyNotes
func GetItemBuyNotes() *ItemBuyNotes {
	return poolItemBuyNotes.Get().(*ItemBuyNotes)
}

// ReleaseItemBuyNotes 释放ItemBuyNotes
func ReleaseItemBuyNotes(v *ItemBuyNotes) {
	v.BuyNotesDetailList = v.BuyNotesDetailList[:0]
	v.NotesTitle = ""
	poolItemBuyNotes.Put(v)
}
