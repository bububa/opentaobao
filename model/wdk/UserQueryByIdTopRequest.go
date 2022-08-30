package wdk

// UserQueryByIdTopRequest 结构体
type UserQueryByIdTopRequest struct {
	// 仓库code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 人员id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
