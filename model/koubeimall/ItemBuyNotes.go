package koubeimall

// ItemBuyNotes 结构体
type ItemBuyNotes struct {
	// 列表
	BuyNotesDetailList []ItemBuyNotesDetail `json:"buy_notes_detail_list,omitempty" xml:"buy_notes_detail_list>item_buy_notes_detail,omitempty"`
	// 购买须知标题
	NotesTitle string `json:"notes_title,omitempty" xml:"notes_title,omitempty"`
}
