package wdk

// DpsScanContainerMtopRequest 结构体
type DpsScanContainerMtopRequest struct {
	// 明细列表
	DetailIdList []int64 `json:"detail_id_list,omitempty" xml:"detail_id_list>int64,omitempty"`
	// 用户
	UserAccount string `json:"user_account,omitempty" xml:"user_account,omitempty"`
	// 容器号
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}
