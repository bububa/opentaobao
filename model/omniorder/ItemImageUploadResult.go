package omniorder

// ItemImageUploadResult 
type ItemImageUploadResult struct {

    // imgId
    
    ImgId   int64 `json:"img_id,omitempty" xml:"img_id,omitempty"`
    

    // itemId
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // url
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    

    // duplicateInfos
    
    DuplicateInfos   []ItemSkuDuplicateInfo `json:"duplicate_infos,omitempty" xml:"duplicate_infos,omitempty"`
    

}
