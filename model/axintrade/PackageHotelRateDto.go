package axintrade

// PackageHotelRateDto 结构体
type PackageHotelRateDto struct {
	// 酒店床型名称
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 酒店rate信息
	HotelRate *HotelRateDto `json:"hotel_rate,omitempty" xml:"hotel_rate,omitempty"`
	// 酒店间夜数
	RoomNightNum int64 `json:"room_night_num,omitempty" xml:"room_night_num,omitempty"`
	// 成人数
	AdultNum int64 `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	// 儿童数
	ChildrenNum int64 `json:"children_num,omitempty" xml:"children_num,omitempty"`
}
