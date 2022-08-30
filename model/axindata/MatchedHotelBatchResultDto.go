package axindata

// MatchedHotelBatchResultDto 结构体
type MatchedHotelBatchResultDto struct {
	// 酒店匹配结果列表
	MatchedHotelResultList []MatchedHotelResultDto `json:"matched_hotel_result_list,omitempty" xml:"matched_hotel_result_list>matched_hotel_result_dto,omitempty"`
}
