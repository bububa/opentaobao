package alihouse

// ProjectSalesTimeDto 结构体
type ProjectSalesTimeDto struct {
	// 预售信息对象列表
	PrePermitDTOList []ProjectSalesTimePrePermitDto `json:"pre_permit_d_t_o_list,omitempty" xml:"pre_permit_d_t_o_list>project_sales_time_pre_permit_dto,omitempty"`
	// 外部楼栋id列表
	OuterTidList []string `json:"outer_tid_list,omitempty" xml:"outer_tid_list>string,omitempty"`
	// 外部项目id
	OuterSid string `json:"outer_sid,omitempty" xml:"outer_sid,omitempty"`
	// 外部楼盘id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 城市名
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 项目名称
	SubjectName string `json:"subject_name,omitempty" xml:"subject_name,omitempty"`
	// 项目价格
	SubjectAvgPrice string `json:"subject_avg_price,omitempty" xml:"subject_avg_price,omitempty"`
	// 装修价格
	DecorationStandardPriceDesc string `json:"decoration_standard_price_desc,omitempty" xml:"decoration_standard_price_desc,omitempty"`
	// 登记开始时间
	RegisterStartDate string `json:"register_start_date,omitempty" xml:"register_start_date,omitempty"`
	// 登记结束时间
	RegisterEndDate string `json:"register_end_date,omitempty" xml:"register_end_date,omitempty"`
	// 认筹开始时间
	RecognizeStartDate string `json:"recognize_start_date,omitempty" xml:"recognize_start_date,omitempty"`
	// 认筹结束时间
	RecognizeEndDate string `json:"recognize_end_date,omitempty" xml:"recognize_end_date,omitempty"`
	// 交资料开始时间
	MaterialStartDate string `json:"material_start_date,omitempty" xml:"material_start_date,omitempty"`
	// 交资料结束时间
	MaterialEndDate string `json:"material_end_date,omitempty" xml:"material_end_date,omitempty"`
	// 公示开始时间
	PublicityStartDate string `json:"publicity_start_date,omitempty" xml:"publicity_start_date,omitempty"`
	// 公示结束时间
	PublicityEndDate string `json:"publicity_end_date,omitempty" xml:"publicity_end_date,omitempty"`
	// 摇号时间
	LotteryDrawDate string `json:"lottery_draw_date,omitempty" xml:"lottery_draw_date,omitempty"`
	// 线下摇号
	OfflineLotteryDrawDesc string `json:"offline_lottery_draw_desc,omitempty" xml:"offline_lottery_draw_desc,omitempty"`
	// 开盘时间
	OpenDate string `json:"open_date,omitempty" xml:"open_date,omitempty"`
	// 选房开始时间
	SelectionStartDate string `json:"selection_start_date,omitempty" xml:"selection_start_date,omitempty"`
	// 选房结束时间
	SelectionEndDate string `json:"selection_end_date,omitempty" xml:"selection_end_date,omitempty"`
	// 线下选房
	OfflineRoomSelectionDesc string `json:"offline_room_selection_desc,omitempty" xml:"offline_room_selection_desc,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 预计交房时间
	DeveloperDueTime string `json:"developer_due_time,omitempty" xml:"developer_due_time,omitempty"`
	// 冻资金额
	FrozenPrice string `json:"frozen_price,omitempty" xml:"frozen_price,omitempty"`
	// 冻资银行
	FrozenBank string `json:"frozen_bank,omitempty" xml:"frozen_bank,omitempty"`
	// 公示方案
	PublicPlan string `json:"public_plan,omitempty" xml:"public_plan,omitempty"`
	// 开盘户型范围
	OpeningLayoutRange string `json:"opening_layout_range,omitempty" xml:"opening_layout_range,omitempty"`
	// 户型面积区间
	LayoutAreaRange string `json:"layout_area_range,omitempty" xml:"layout_area_range,omitempty"`
	// 开盘楼栋范围
	OpeningBuildingRange string `json:"opening_building_range,omitempty" xml:"opening_building_range,omitempty"`
	// 城市编码
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 有效状态 0-无效 1-有效
	IsValid int64 `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
	// 是否具有登记环节 0(否) 1(是)
	IsRegister int64 `json:"is_register,omitempty" xml:"is_register,omitempty"`
	// 是否具有认筹环节 0(否) 1(是)
	IsRecognize int64 `json:"is_recognize,omitempty" xml:"is_recognize,omitempty"`
	// 是否具有交资料环节 0(否) 1(是)
	IsMaterial int64 `json:"is_material,omitempty" xml:"is_material,omitempty"`
	// 是否具有公示环节 0(否) 1(是)
	IsPublicity int64 `json:"is_publicity,omitempty" xml:"is_publicity,omitempty"`
	// 是否具有摇号环节 0(否) 1(是)
	IsLottery int64 `json:"is_lottery,omitempty" xml:"is_lottery,omitempty"`
	// 是否具有开盘环节 0(否) 1(是)
	IsOpen int64 `json:"is_open,omitempty" xml:"is_open,omitempty"`
	// 是否模糊显示 0-否 1-是
	IsBlurOpen int64 `json:"is_blur_open,omitempty" xml:"is_blur_open,omitempty"`
	// 是否具有选房环节 0(否) 1(是)
	IsSelection int64 `json:"is_selection,omitempty" xml:"is_selection,omitempty"`
	// 摇号比例信息
	LotteryRatioDTO *LotteryRatioDto `json:"lottery_ratio_d_t_o,omitempty" xml:"lottery_ratio_d_t_o,omitempty"`
}
