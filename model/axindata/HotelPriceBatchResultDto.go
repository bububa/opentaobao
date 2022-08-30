package axindata

// HotelPriceBatchResultDto 结构体
type HotelPriceBatchResultDto struct {
	// 酒店的报价结果列表
	HotelPriceDtoList []HotelPriceBatchDto `json:"hotel_price_dto_list,omitempty" xml:"hotel_price_dto_list>hotel_price_batch_dto,omitempty"`
}
