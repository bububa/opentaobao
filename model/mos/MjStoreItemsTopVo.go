package mos

// MjStoreItemsTopVo 
type MjStoreItemsTopVo struct {

    // 商品列表
    ItemList   []MjItemTopVo `json:"item_list,omitempty"`

    // 专柜id
    StoreId   int64 `json:"store_id,omitempty"`

    // 专柜名称
    StoreName   string `json:"store_name,omitempty"`

    // uuid
    Uuid   string `json:"uuid,omitempty"`

    // bucketId
    BucketId   string `json:"bucket_id,omitempty"`

}
