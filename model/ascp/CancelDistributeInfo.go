package ascp

// CancelDistributeInfo 结构体
type CancelDistributeInfo struct {
	// 要取消铺货的分销商ID列表,如果 cancelAll = true, 则不需要设置此字段,如果 cancelAll = false， 则此字段必填
	DistributorShopUserIdList []int64 `json:"distributor_shop_user_id_list,omitempty" xml:"distributor_shop_user_id_list>int64,omitempty"`
	// 【必传】要取消的商品itemId
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 要取消的商品skuId, 如果没有则不需要传
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 【必传】取消所有铺货,如果设置为false，则需要设置  distributorShopUserIdList
	CancelAll bool `json:"cancel_all,omitempty" xml:"cancel_all,omitempty"`
}
