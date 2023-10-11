package ascp

// WarehouseCooperationConfirmResultDetail 结构体
type WarehouseCooperationConfirmResultDetail struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 合作店铺仓code
	CooperationWhCode string `json:"cooperation_wh_code,omitempty" xml:"cooperation_wh_code,omitempty"`
}
