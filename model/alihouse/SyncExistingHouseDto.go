package alihouse

// SyncExistingHouseDto 结构体
type SyncExistingHouseDto struct {
	// 户型图片
	LayoutPictures []string `json:"layout_pictures,omitempty" xml:"layout_pictures>string,omitempty"`
	// 室内图
	InnerPictures []string `json:"inner_pictures,omitempty" xml:"inner_pictures>string,omitempty"`
	// 房源标签
	TagCodes []string `json:"tag_codes,omitempty" xml:"tag_codes>string,omitempty"`
	// 图片dto
	Pictures []ProjectPictureDto `json:"pictures,omitempty" xml:"pictures>project_picture_dto,omitempty"`
	// 货品商品关联DTo
	Extend []GoodsRelationDto `json:"extend,omitempty" xml:"extend>goods_relation_dto,omitempty"`
	// 货品商品关联DTo
	RelationCargos []GoodsRelationDto `json:"relation_cargos,omitempty" xml:"relation_cargos>goods_relation_dto,omitempty"`
	// 首录时间
	GatherTime string `json:"gather_time,omitempty" xml:"gather_time,omitempty"`
	// 视频链接
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 价格 单位万元
	ShowPrice string `json:"show_price,omitempty" xml:"show_price,omitempty"`
	// 建筑年代
	ConstructionTime string `json:"construction_time,omitempty" xml:"construction_time,omitempty"`
	// 房源上下架原因
	HouseReason string `json:"house_reason,omitempty" xml:"house_reason,omitempty"`
	// 看房时间
	LookTime string `json:"look_time,omitempty" xml:"look_time,omitempty"`
	// 税费描述
	FeeDesc string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	// 服务描述
	ServiceDesc string `json:"service_desc,omitempty" xml:"service_desc,omitempty"`
	// 小区描述
	CommunityDesc string `json:"community_desc,omitempty" xml:"community_desc,omitempty"`
	// 户型描述
	LayoutDesc string `json:"layout_desc,omitempty" xml:"layout_desc,omitempty"`
	// 业主心态
	Mentality string `json:"mentality,omitempty" xml:"mentality,omitempty"`
	// 房源介绍
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 房号
	RoomNo string `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// 单元
	UnitNo string `json:"unit_no,omitempty" xml:"unit_no,omitempty"`
	// 楼栋
	BuildingNo string `json:"building_no,omitempty" xml:"building_no,omitempty"`
	// 建筑面积
	BuildingArea string `json:"building_area,omitempty" xml:"building_area,omitempty"`
	// 套内面积
	InsideArea string `json:"inside_area,omitempty" xml:"inside_area,omitempty"`
	// 房源标题
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 房源地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 房源E码
	Ecode string `json:"ecode,omitempty" xml:"ecode,omitempty"`
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 视频封面图
	VideoCoverImage string `json:"video_cover_image,omitempty" xml:"video_cover_image,omitempty"`
	// 小区名
	CommunityName string `json:"community_name,omitempty" xml:"community_name,omitempty"`
	// 外部户型id
	OuterLayoutsId string `json:"outer_layouts_id,omitempty" xml:"outer_layouts_id,omitempty"`
	// 房屋设施
	ServiceFacilities string `json:"service_facilities,omitempty" xml:"service_facilities,omitempty"`
	// 租金
	RentPrice string `json:"rent_price,omitempty" xml:"rent_price,omitempty"`
	// 外部房屋ID
	OuterHouseBaseId string `json:"outer_house_base_id,omitempty" xml:"outer_house_base_id,omitempty"`
	// 单价，XX元/m²，整数
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 房源委托信息
	Entrust *SyncExistingHouseEntrustDto `json:"entrust,omitempty" xml:"entrust,omitempty"`
	// 是否展示 1展示 0不展示
	IsShow int64 `json:"is_show,omitempty" xml:"is_show,omitempty"`
	// 产权年限
	PropertyAge int64 `json:"property_age,omitempty" xml:"property_age,omitempty"`
	// 是否唯一住房 0不是 1是
	IsOnlyHouse int64 `json:"is_only_house,omitempty" xml:"is_only_house,omitempty"`
	// 是否有电梯 0 没有 1有
	IsElevator int64 `json:"is_elevator,omitempty" xml:"is_elevator,omitempty"`
	// 房源状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 物业性质
	PropertyType int64 `json:"property_type,omitempty" xml:"property_type,omitempty"`
	// 总楼层
	TotalFloor int64 `json:"total_floor,omitempty" xml:"total_floor,omitempty"`
	// 当前楼层
	Floor int64 `json:"floor,omitempty" xml:"floor,omitempty"`
	// 楼层位置
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
	// 阳台数
	Balcony int64 `json:"balcony,omitempty" xml:"balcony,omitempty"`
	// 厨房数
	Kitchen int64 `json:"kitchen,omitempty" xml:"kitchen,omitempty"`
	// 厕所数
	Toilet int64 `json:"toilet,omitempty" xml:"toilet,omitempty"`
	// 客厅数
	Hall int64 `json:"hall,omitempty" xml:"hall,omitempty"`
	// 室数
	Room int64 `json:"room,omitempty" xml:"room,omitempty"`
	// 朝向
	Direction int64 `json:"direction,omitempty" xml:"direction,omitempty"`
	// 装修标准
	Fitment int64 `json:"fitment,omitempty" xml:"fitment,omitempty"`
	// 房屋现状
	Situation int64 `json:"situation,omitempty" xml:"situation,omitempty"`
	// 土地用途
	LandUse int64 `json:"land_use,omitempty" xml:"land_use,omitempty"`
	// 产权性质
	PropertyRights int64 `json:"property_rights,omitempty" xml:"property_rights,omitempty"`
	// 公私域区分 1公域 2私域
	Scene int64 `json:"scene,omitempty" xml:"scene,omitempty"`
	// 0-正常房源 1-临时房源（交易专用）
	HouseType int64 `json:"house_type,omitempty" xml:"house_type,omitempty"`
	// 是否存在室内图封面,1是，2否
	IsCover int64 `json:"is_cover,omitempty" xml:"is_cover,omitempty"`
	// 是否需要异步
	IsAsync int64 `json:"is_async,omitempty" xml:"is_async,omitempty"`
	// 出租状态
	RentStatus int64 `json:"rent_status,omitempty" xml:"rent_status,omitempty"`
	// 隔断类型
	PartitionType int64 `json:"partition_type,omitempty" xml:"partition_type,omitempty"`
	// 窗类型
	WindowType int64 `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 是否有独立阳台
	HasIndependentBalcony int64 `json:"has_independent_balcony,omitempty" xml:"has_independent_balcony,omitempty"`
	// 是否有独立卫生间
	HasIndependentToilet int64 `json:"has_independent_toilet,omitempty" xml:"has_independent_toilet,omitempty"`
	// 租房业务模式
	HouseModel int64 `json:"house_model,omitempty" xml:"house_model,omitempty"`
	// 房源分类
	HouseCategory int64 `json:"house_category,omitempty" xml:"house_category,omitempty"`
	// 房间类型
	RoomType int64 `json:"room_type,omitempty" xml:"room_type,omitempty"`
	// 业务类型，1-新房，2-二手房，3-租房，默认为2
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 0默认不是测试，1是测试数据
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 货类型 6-楼栋  7-户型  8-房源
	CargoType int64 `json:"cargo_type,omitempty" xml:"cargo_type,omitempty"`
	// 货的销售状态 0-待定 1-待售 2-在售 3-售罄 4-停售
	SalesStatus int64 `json:"sales_status,omitempty" xml:"sales_status,omitempty"`
	// 货的商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
