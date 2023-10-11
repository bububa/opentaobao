package ascp

// CooperationWarehouses 结构体
type CooperationWarehouses struct {
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 合作店铺仓code
	CooperationWhCode string `json:"cooperation_wh_code,omitempty" xml:"cooperation_wh_code,omitempty"`
}
