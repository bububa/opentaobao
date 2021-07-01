package mos

// MjStoreItemsTopVo 结构体
type MjStoreItemsTopVo struct {
	// 商品列表
	ItemList []MjItemTopVo `json:"item_list,omitempty" xml:"item_list>mj_item_top_vo,omitempty"`
	// 专柜id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 专柜名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// bucketId
	BucketId string `json:"bucket_id,omitempty" xml:"bucket_id,omitempty"`
}
