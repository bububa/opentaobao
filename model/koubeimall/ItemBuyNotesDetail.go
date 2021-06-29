package koubeimall

// ItemBuyNotesDetail 
type ItemBuyNotesDetail struct {
    // 须知标题
    NotesDetailTitle   string `json:"notes_detail_title,omitempty" xml:"notes_detail_title,omitempty"`
    // 须知明细
    NotesDetails   []string `json:"notes_details,omitempty" xml:"notes_details>string,omitempty"`
}
