package alihouse

// ProjectPictureDto 结构体
type ProjectPictureDto struct {
	// 是否删除
	IsDeleted string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 图片说明
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 图片URL全路径
	PicData string `json:"pic_data,omitempty" xml:"pic_data,omitempty"`
	// 外部房源ID
	OuterHouseId string `json:"outer_house_id,omitempty" xml:"outer_house_id,omitempty"`
	// 外部房屋ID
	OuterHouseBaseId string `json:"outer_house_base_id,omitempty" xml:"outer_house_base_id,omitempty"`
	// 外部户型ID
	OuterLayoutId string `json:"outer_layout_id,omitempty" xml:"outer_layout_id,omitempty"`
	// 外部图片ID
	OuterPictureId string `json:"outer_picture_id,omitempty" xml:"outer_picture_id,omitempty"`
	// 外部id包含楼盘、小区
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 图片名称
	PicName string `json:"pic_name,omitempty" xml:"pic_name,omitempty"`
	// 11111
	OuterLayoutTid string `json:"outer_layout_tid,omitempty" xml:"outer_layout_tid,omitempty"`
	// 22222
	OuterBuildingId string `json:"outer_building_id,omitempty" xml:"outer_building_id,omitempty"`
	// 1111
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 是否封面图
	IsCover int64 `json:"is_cover,omitempty" xml:"is_cover,omitempty"`
	// 是否焦点图
	IsFocus int64 `json:"is_focus,omitempty" xml:"is_focus,omitempty"`
	// 数据类型1-楼盘，2-小区，3-租房
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// 图片类型
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 图片顺序
	OrderTag int64 `json:"order_tag,omitempty" xml:"order_tag,omitempty"`
}
