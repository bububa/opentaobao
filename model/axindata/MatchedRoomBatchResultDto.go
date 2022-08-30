package axindata

// MatchedRoomBatchResultDto 结构体
type MatchedRoomBatchResultDto struct {
	// 每个房型匹配结果封装对象
	MatchedRoomResultList []MatchedRoomResultDto `json:"matched_room_result_list,omitempty" xml:"matched_room_result_list>matched_room_result_dto,omitempty"`
}
