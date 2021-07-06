package ascpchannel

// MerchantInventoryQuery 结构体
type MerchantInventoryQuery struct {
	// 供应链货主id list 单次<=20
	ScItemIds []int64 `json:"sc_item_ids,omitempty" xml:"sc_item_ids>int64,omitempty"`
	// 供应链中台商家仓code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 供应链中台供应商id
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 供应链中台物流货主id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
