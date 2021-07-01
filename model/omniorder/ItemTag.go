package omniorder

// ItemTag 结构体
type ItemTag struct {
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// tagType
	TagType string `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
}
