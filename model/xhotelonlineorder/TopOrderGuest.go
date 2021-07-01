package xhotelonlineorder

// TopOrderGuest 结构体
type TopOrderGuest struct {
	// 住客类型
	CustomerType int64 `json:"customer_type,omitempty" xml:"customer_type,omitempty"`
	// 住客编号
	PersonNo int64 `json:"person_no,omitempty" xml:"person_no,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 年龄
	Age int64 `json:"age,omitempty" xml:"age,omitempty"`
	// 房间编号
	RoomNo int64 `json:"room_no,omitempty" xml:"room_no,omitempty"`
}
