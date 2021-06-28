package mos

// MjStoreItemsTopVo 
/* model for simplify = false
type MjStoreItemsTopVo struct {

    // 商品列表
    
    ItemList  struct {
        MjItemTopVo  []MjItemTopVo `json:"mj_item_top_vo,omitempty"`
    } `json:"item_list,omitempty"`
    

    // 专柜id
    
    StoreId   int64 `json:"store_id,omitempty"`
    

    // 专柜名称
    
    StoreName   string `json:"store_name,omitempty"`
    

    // uuid
    
    Uuid   string `json:"uuid,omitempty"`
    

    // bucketId
    
    BucketId   string `json:"bucket_id,omitempty"`
    

}
*/

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
