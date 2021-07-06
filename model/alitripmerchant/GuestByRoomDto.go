package alitripmerchant

// GuestByRoomDto 结构体
type GuestByRoomDto struct {
	// 儿童年龄数组
	ChildAges []int64 `json:"child_ages,omitempty" xml:"child_ages>int64,omitempty"`
	// 入住人名
	ContactFirstName string `json:"contact_first_name,omitempty" xml:"contact_first_name,omitempty"`
	// 入住人姓
	ContactLastName string `json:"contact_last_name,omitempty" xml:"contact_last_name,omitempty"`
	// 当前房间的总人数
	RoomerNumber int64 `json:"roomer_number,omitempty" xml:"roomer_number,omitempty"`
	// 成人数
	AdultRoomerNumber int64 `json:"adult_roomer_number,omitempty" xml:"adult_roomer_number,omitempty"`
	// 儿童数
	ChildRoomerNumber int64 `json:"child_roomer_number,omitempty" xml:"child_roomer_number,omitempty"`
}
