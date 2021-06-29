package xhotelonlineorder

// GuestInfo 
type GuestInfo struct {
    // 入住人在这个房间的序号（从1开始）
    PersonPos   string `json:"person_pos,omitempty" xml:"person_pos,omitempty"`
    // 房间序号（从1开始）
    RoomPos   string `json:"room_pos,omitempty" xml:"room_pos,omitempty"`
    // 入住人姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
