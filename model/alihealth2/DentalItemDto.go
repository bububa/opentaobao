package alihealth2

// DentalItemDto 结构体
type DentalItemDto struct {
	// itemName
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
