package flight

// PolicyResponseDto 结构体
type PolicyResponseDto struct {
	// 运价渠道
	FareSources []string `json:"fare_sources,omitempty" xml:"fare_sources>string,omitempty"`
	// 行程限制
	Flights []FlightDto `json:"flights,omitempty" xml:"flights>flight_dto,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 政策代码
	PolicyCode string `json:"policy_code,omitempty" xml:"policy_code,omitempty"`
	// 共享航司二字码
	CodeShareAirline string `json:"code_share_airline,omitempty" xml:"code_share_airline,omitempty"`
	// 降落机场
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 航空公司二字码
	Airline string `json:"airline,omitempty" xml:"airline,omitempty"`
	// 大客户编码
	AccountCode string `json:"account_code,omitempty" xml:"account_code,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 到达机场
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// farebasis
	FareBasis string `json:"fare_basis,omitempty" xml:"fare_basis,omitempty"`
	// 错误代码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 商家配置号
	OfficeNo string `json:"office_no,omitempty" xml:"office_no,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 下浮
	Down int64 `json:"down,omitempty" xml:"down,omitempty"`
	// 紧张阈值
	CreatePnrLimit int64 `json:"create_pnr_limit,omitempty" xml:"create_pnr_limit,omitempty"`
	// 支持共享航班
	CodeShare int64 `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// 价格控制
	Price *PriceDto `json:"price,omitempty" xml:"price,omitempty"`
	// 上浮
	Up int64 `json:"up,omitempty" xml:"up,omitempty"`
	// 库存
	Stock *StockDto `json:"stock,omitempty" xml:"stock,omitempty"`
	// 政策来源
	PolicySource int64 `json:"policy_source,omitempty" xml:"policy_source,omitempty"`
	// 浮动单位
	FloatUnit int64 `json:"float_unit,omitempty" xml:"float_unit,omitempty"`
	// 票面价（元）
	FarePrice int64 `json:"fare_price,omitempty" xml:"fare_price,omitempty"`
	// pata
	Pata int64 `json:"pata,omitempty" xml:"pata,omitempty"`
	// 是否订位
	CreatePnr int64 `json:"create_pnr,omitempty" xml:"create_pnr,omitempty"`
	// 行程类型
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// 销售限制
	Sale *SaleDto `json:"sale,omitempty" xml:"sale,omitempty"`
	// 乘客限制
	Passenger *PassengerDto `json:"passenger,omitempty" xml:"passenger,omitempty"`
	// 政策类型
	PolicyType int64 `json:"policy_type,omitempty" xml:"policy_type,omitempty"`
	// 舱位条件
	IsWhite int64 `json:"is_white,omitempty" xml:"is_white,omitempty"`
	// 政策状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
