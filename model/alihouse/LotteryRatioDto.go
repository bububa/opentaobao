package alihouse

// LotteryRatioDto 结构体
type LotteryRatioDto struct {
	// 人才入围条件
	TalentShortlisted string `json:"talent_shortlisted,omitempty" xml:"talent_shortlisted,omitempty"`
	// 无房入围条件
	NoRoomShortlisted string `json:"no_room_shortlisted,omitempty" xml:"no_room_shortlisted,omitempty"`
	// 有房入围条件
	YesRoomShortlisted string `json:"yes_room_shortlisted,omitempty" xml:"yes_room_shortlisted,omitempty"`
	// 总体中签率
	TotalSuccessRate string `json:"total_success_rate,omitempty" xml:"total_success_rate,omitempty"`
	// 人才中签率
	TalentSuccessRate string `json:"talent_success_rate,omitempty" xml:"talent_success_rate,omitempty"`
	// 无房中签率
	NoRoomSuccessRate string `json:"no_room_success_rate,omitempty" xml:"no_room_success_rate,omitempty"`
	// 有房中签率
	YesRoomSuccessRate string `json:"yes_room_success_rate,omitempty" xml:"yes_room_success_rate,omitempty"`
	// 摇号比例类型 1-廊内 2-廊外
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 有无人才优惠 (0-无 1-有)
	HasExistenceTalent int64 `json:"has_existence_talent,omitempty" xml:"has_existence_talent,omitempty"`
	// 有无入围条件 (0-无 1-有)
	HasShortlisted int64 `json:"has_shortlisted,omitempty" xml:"has_shortlisted,omitempty"`
	// 全部房源套数
	TotalHouseNumber int64 `json:"total_house_number,omitempty" xml:"total_house_number,omitempty"`
	// 人才房源套数
	TalentHouseNumber int64 `json:"talent_house_number,omitempty" xml:"talent_house_number,omitempty"`
	// 无房房源套数
	NoRoomHouseNumber int64 `json:"no_room_house_number,omitempty" xml:"no_room_house_number,omitempty"`
	// 有房房源套数
	YesRoomHouseNumber int64 `json:"yes_room_house_number,omitempty" xml:"yes_room_house_number,omitempty"`
	// 全部入围人数
	TotalShortlistedNumber int64 `json:"total_shortlisted_number,omitempty" xml:"total_shortlisted_number,omitempty"`
	// 人才入围人数
	TalentShortlistedNumber int64 `json:"talent_shortlisted_number,omitempty" xml:"talent_shortlisted_number,omitempty"`
	// 无房入围人数
	NoRoomShortlistedNumber int64 `json:"no_room_shortlisted_number,omitempty" xml:"no_room_shortlisted_number,omitempty"`
	// 有房入围人数
	YesRoomShortlistedNumber int64 `json:"yes_room_shortlisted_number,omitempty" xml:"yes_room_shortlisted_number,omitempty"`
}
