package iotticket

// RepairmanInfo 结构体
type RepairmanInfo struct {
	// 上门维修人员编号
	RepairmanId string `json:"repairman_id,omitempty" xml:"repairman_id,omitempty"`
	// 上门维修人员名称
	RepairmanName string `json:"repairman_name,omitempty" xml:"repairman_name,omitempty"`
	// 上门维修人员联系方式
	RepairmanPhone string `json:"repairman_phone,omitempty" xml:"repairman_phone,omitempty"`
	// 上门时间
	AppointDate string `json:"appoint_date,omitempty" xml:"appoint_date,omitempty"`
	// 上门维修地址
	VisitAddress string `json:"visit_address,omitempty" xml:"visit_address,omitempty"`
}
