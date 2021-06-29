package koubeimall

// ItemBuyNotes 
type ItemBuyNotes struct {

    // 购买须知标题
    
    NotesTitle   string `json:"notes_title,omitempty" xml:"notes_title,omitempty"`
    

    // 列表
    
    BuyNotesDetailList   []ItemBuyNotesDetail `json:"buy_notes_detail_list,omitempty" xml:"buy_notes_detail_list,omitempty"`
    

}
