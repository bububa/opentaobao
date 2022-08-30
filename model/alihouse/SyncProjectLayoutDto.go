package alihouse

// SyncProjectLayoutDto 结构体
type SyncProjectLayoutDto struct {
	// 绑定的品
	Extend []GoodsRelationDto `json:"extend,omitempty" xml:"extend>goods_relation_dto,omitempty"`
	// 绑定的货
	RelationCargos []GoodsRelationDto `json:"relation_cargos,omitempty" xml:"relation_cargos>goods_relation_dto,omitempty"`
	// 朝向
	Orientation string `json:"orientation,omitempty" xml:"orientation,omitempty"`
	// 户型名称
	LayoutName string `json:"layout_name,omitempty" xml:"layout_name,omitempty"`
	// 外部户型ID
	OuterLayoutId string `json:"outer_layout_id,omitempty" xml:"outer_layout_id,omitempty"`
	// 户型图
	LayoutImages string `json:"layout_images,omitempty" xml:"layout_images,omitempty"`
	// 均价
	AvgPrice string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
	// 总价
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 套内面积
	InsideArea string `json:"inside_area,omitempty" xml:"inside_area,omitempty"`
	// 建筑面积
	ConstructionArea string `json:"construction_area,omitempty" xml:"construction_area,omitempty"`
	// 卫
	Bathroom string `json:"bathroom,omitempty" xml:"bathroom,omitempty"`
	// 厅
	Hall string `json:"hall,omitempty" xml:"hall,omitempty"`
	// 室
	Room string `json:"room,omitempty" xml:"room,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 销售状态
	SalesStatus string `json:"sales_status,omitempty" xml:"sales_status,omitempty"`
	// 户型标签
	LayoutLabel string `json:"layout_label,omitempty" xml:"layout_label,omitempty"`
	// 厨房数量
	Kitchen string `json:"kitchen,omitempty" xml:"kitchen,omitempty"`
	// 服务设施 逗号隔开
	ServiceFacility string `json:"service_facility,omitempty" xml:"service_facility,omitempty"`
	// 楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 是否删除 1 是 0 否
	IsDeleted string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部货ID
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 是否是主力户型
	IsMainLayout int64 `json:"is_main_layout,omitempty" xml:"is_main_layout,omitempty"`
	// 是否有开放厨房 1-是 0-否
	OpenKitchen int64 `json:"open_kitchen,omitempty" xml:"open_kitchen,omitempty"`
	// 业务类型 3-租房
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 物业类型
	EstateType int64 `json:"estate_type,omitempty" xml:"estate_type,omitempty"`
}
