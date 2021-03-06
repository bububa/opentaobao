package omniorder

// ItemImageUploadResult 结构体
type ItemImageUploadResult struct {
	// duplicateInfos
	DuplicateInfos []ItemSkuDuplicateInfo `json:"duplicate_infos,omitempty" xml:"duplicate_infos>item_sku_duplicate_info,omitempty"`
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// imgId
	ImgId int64 `json:"img_id,omitempty" xml:"img_id,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
