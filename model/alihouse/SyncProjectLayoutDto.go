package alihouse

import (
	"sync"
)

// SyncProjectLayoutDto 结构体
type SyncProjectLayoutDto struct {
	// 绑定的品
	Extend []GoodsRelationDto `json:"extend,omitempty" xml:"extend>goods_relation_dto,omitempty"`
	// 绑定的货
	RelationCargos []GoodsRelationDto `json:"relation_cargos,omitempty" xml:"relation_cargos>goods_relation_dto,omitempty"`
	// 1,2
	OuterConsultantIds []string `json:"outer_consultant_ids,omitempty" xml:"outer_consultant_ids>string,omitempty"`
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
	// 视频封面图
	VideoCoverImg string `json:"video_cover_img,omitempty" xml:"video_cover_img,omitempty"`
	// 视频URL
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// eCode
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 外部经纪人ID
	OuterBrokerId string `json:"outer_broker_id,omitempty" xml:"outer_broker_id,omitempty"`
	// 1栋
	BuildingDescribe string `json:"building_describe,omitempty" xml:"building_describe,omitempty"`
	// 1
	Balcony string `json:"balcony,omitempty" xml:"balcony,omitempty"`
	// 1,2
	BelongBuildings string `json:"belong_buildings,omitempty" xml:"belong_buildings,omitempty"`
	// 1,2
	BelongUnits string `json:"belong_units,omitempty" xml:"belong_units,omitempty"`
	// 1,2
	BelongRooms string `json:"belong_rooms,omitempty" xml:"belong_rooms,omitempty"`
	// 1,2
	RoomNumber string `json:"room_number,omitempty" xml:"room_number,omitempty"`
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
	// 装修标准
	DecorationStandard int64 `json:"decoration_standard,omitempty" xml:"decoration_standard,omitempty"`
}

var poolSyncProjectLayoutDto = sync.Pool{
	New: func() any {
		return new(SyncProjectLayoutDto)
	},
}

// GetSyncProjectLayoutDto() 从对象池中获取SyncProjectLayoutDto
func GetSyncProjectLayoutDto() *SyncProjectLayoutDto {
	return poolSyncProjectLayoutDto.Get().(*SyncProjectLayoutDto)
}

// ReleaseSyncProjectLayoutDto 释放SyncProjectLayoutDto
func ReleaseSyncProjectLayoutDto(v *SyncProjectLayoutDto) {
	v.Extend = v.Extend[:0]
	v.RelationCargos = v.RelationCargos[:0]
	v.OuterConsultantIds = v.OuterConsultantIds[:0]
	v.Orientation = ""
	v.LayoutName = ""
	v.OuterLayoutId = ""
	v.LayoutImages = ""
	v.AvgPrice = ""
	v.TotalPrice = ""
	v.InsideArea = ""
	v.ConstructionArea = ""
	v.Bathroom = ""
	v.Hall = ""
	v.Room = ""
	v.Description = ""
	v.SalesStatus = ""
	v.LayoutLabel = ""
	v.Kitchen = ""
	v.ServiceFacility = ""
	v.OuterId = ""
	v.IsDeleted = ""
	v.OuterStoreId = ""
	v.OuterTid = ""
	v.VideoCoverImg = ""
	v.VideoUrl = ""
	v.ECode = ""
	v.OuterBrokerId = ""
	v.BuildingDescribe = ""
	v.Balcony = ""
	v.BelongBuildings = ""
	v.BelongUnits = ""
	v.BelongRooms = ""
	v.RoomNumber = ""
	v.IsMainLayout = 0
	v.OpenKitchen = 0
	v.SourceType = 0
	v.ItemId = 0
	v.Type = 0
	v.EstateType = 0
	v.DecorationStandard = 0
	poolSyncProjectLayoutDto.Put(v)
}
