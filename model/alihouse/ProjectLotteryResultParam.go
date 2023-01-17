package alihouse

// ProjectLotteryResultParam 结构体
type ProjectLotteryResultParam struct {
	// 摇号结果DTO
	ProjectLotteryResultDTOList []ProjectLotteryResultDto `json:"project_lottery_result_d_t_o_list,omitempty" xml:"project_lottery_result_d_t_o_list>project_lottery_result_dto,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部楼盘ID
	OuterProjectId string `json:"outer_project_id,omitempty" xml:"outer_project_id,omitempty"`
	// 外部分期ID
	OuterSalesTimeId string `json:"outer_sales_time_id,omitempty" xml:"outer_sales_time_id,omitempty"`
	// 外部摇号结果ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 类型 1意向登记表 2摇号结果表
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 外部版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 是否删除
	DeleteFlag bool `json:"delete_flag,omitempty" xml:"delete_flag,omitempty"`
}
