package ascp

// CooperationWarehouse 结构体
type CooperationWarehouse struct {
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 合作店铺仓code
	CooperationWhCode string `json:"cooperation_wh_code,omitempty" xml:"cooperation_wh_code,omitempty"`
	// 店铺sellerid（可空）
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 合作状态： 1：合作待确认 2：合作中 3：已拒绝合作 4：商家取消合作 5：服务商取消合作
	CooperationStatus int64 `json:"cooperation_status,omitempty" xml:"cooperation_status,omitempty"`
}
