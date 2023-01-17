package alihouse

// BuildingDto 结构体
type BuildingDto struct {
	// 货下挂的其他品列表
	Extend []GoodsRelationDto `json:"extend,omitempty" xml:"extend>goods_relation_dto,omitempty"`
	// 货下挂的其他货列表
	RelationCargos []GoodsRelationDto `json:"relation_cargos,omitempty" xml:"relation_cargos>goods_relation_dto,omitempty"`
	// 楼栋E码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 楼栋标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 开盘时间 格式：(yyyy-MM或yyyy-MM-dd)
	DeveloperOpeningTime string `json:"developer_opening_time,omitempty" xml:"developer_opening_time,omitempty"`
	// 交付时间 格式：(yyyy-MM或yyyy-MM-dd)
	DeveloperDueTime string `json:"developer_due_time,omitempty" xml:"developer_due_time,omitempty"`
	// 外部货-楼栋id（外部唯一识别码）
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 外部私域楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 楼栋号
	BuildingNo string `json:"building_no,omitempty" xml:"building_no,omitempty"`
	// 几梯（单层） �&gt;=0
	ElevatorNo int64 `json:"elevator_no,omitempty" xml:"elevator_no,omitempty"`
	// 几户（单层） �&gt;0
	HouseholdNo int64 `json:"household_no,omitempty" xml:"household_no,omitempty"`
	// 单元数 &gt;0
	Units int64 `json:"units,omitempty" xml:"units,omitempty"`
	// 层数 �&gt;0
	Floors int64 `json:"floors,omitempty" xml:"floors,omitempty"`
	// 户数 �&gt;0
	Rooms int64 `json:"rooms,omitempty" xml:"rooms,omitempty"`
	// 货商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货的销售状态 0-待定 1-待售 2-在售 3-售罄 4-停售
	SalesStatus int64 `json:"sales_status,omitempty" xml:"sales_status,omitempty"`
	// 类型 1-安心置业 2-特价房 3-0元购  4-大额电商券 5-认购商品 6-楼栋  7-户型  8-房源
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 是否为货 0-非货 1-货
	IsCargo int64 `json:"is_cargo,omitempty" xml:"is_cargo,omitempty"`
	// 是否有电梯 0-否 1-是
	IsElevator int64 `json:"is_elevator,omitempty" xml:"is_elevator,omitempty"`
}
