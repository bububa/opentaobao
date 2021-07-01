package xhotelonlineorder

// XOrderGuest 结构体
type XOrderGuest struct {
	// 房间序号
	RoomPos int64 `json:"room_pos,omitempty" xml:"room_pos,omitempty"`
	// 入住人序号
	PersonPos int64 `json:"person_pos,omitempty" xml:"person_pos,omitempty"`
	// 入住人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
