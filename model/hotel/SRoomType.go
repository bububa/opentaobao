package hotel

// SRoomType 结构体
type SRoomType struct {
	// 匹配的标准房型
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// 房型名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 宽带服务<br/>"0","有线上网(免费),<br/>"1","有线上网(无)",<br/>"2","有线上网(收费)",<br/>"3","有线上网(部分有且免费)",<br/>"4","有线上网(部分有且收费)"
	Internet string `json:"internet,omitempty" xml:"internet,omitempty"`
	// shid
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// pic_url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// facility
	Facility string `json:"facility,omitempty" xml:"facility,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 扩展字段
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 窗型，枚举类型<br/>0, 无窗,<br/>1, 有窗;
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 床型。json格式：[{"bedType":"大床","bedSize":"1.5m"},{"bedType":"双床","bedSize":"1.2m"}]
	Bed string `json:"bed,omitempty" xml:"bed,omitempty"`
	// 状态。0:正常;-1:删除
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
