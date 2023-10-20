package flight

import (
	"sync"
)

// PriceDto 结构体
type PriceDto struct {
	// 按航段调运价
	FlightPriceValues []FlightPriceDto `json:"flight_price_values,omitempty" xml:"flight_price_values>flight_price_dto,omitempty"`
	// 返点，单位分，如传100表示返点为1%。支持正负数
	Commission int64 `json:"commission,omitempty" xml:"commission,omitempty"`
	// 返点，单位分，如传100表示返点为1%。支持正负数
	ReturnPrice int64 `json:"return_price,omitempty" xml:"return_price,omitempty"`
	// 最小票面价，单位：元
	MinPriceLimit int64 `json:"min_price_limit,omitempty" xml:"min_price_limit,omitempty"`
	// 最大票面价，单位：元
	MaxPriceLimit int64 `json:"max_price_limit,omitempty" xml:"max_price_limit,omitempty"`
	// 固定金额竞价空间，单位分，非必填，非负。录入的数值即为竞价阈值。当固定金额和百分比竞价空间同时有值时，取两者之和做为实际的竞价阈值
	BiddFee int64 `json:"bidd_fee,omitempty" xml:"bidd_fee,omitempty"`
	// 百分比竞价空间，单位分，如传100表示为1%，非必填，支持录入0~100%，根据票面价*百分比竞价空间计算竞价阈值，向下进位到小数点后两位。当固定金额和百分比竞价空间同时有值时，取两者之和做为实际的竞价阈值。
	BiddFeePercent int64 `json:"bidd_fee_percent,omitempty" xml:"bidd_fee_percent,omitempty"`
	// 进位规则
	CarryRule int64 `json:"carry_rule,omitempty" xml:"carry_rule,omitempty"`
	// 最低结算价
	LowestPrice int64 `json:"lowest_price,omitempty" xml:"lowest_price,omitempty"`
	// 票面价算法
	CalFareMethod int64 `json:"cal_fare_method,omitempty" xml:"cal_fare_method,omitempty"`
	// 竞价空间
	BidFee int64 `json:"bid_fee,omitempty" xml:"bid_fee,omitempty"`
	// 竞价方式
	BidMethod int64 `json:"bid_method,omitempty" xml:"bid_method,omitempty"`
	// 儿童价计算方法：选1/2/3/4/5/6/7中一个，不填默认1： 1：儿童成人同价或Y50取低且使用儿童税费 2：儿童成人同价同税 3：儿童自定义价格或Y50取低且使用儿童税费  4：儿童不可销售（仅学生，青年，老年票，小团，学生认证票允许选择） 5：儿童成人同价同税或Y50且使用儿童税费取低 6：儿童成人同价且使用儿童税费 7：使用基准运价FD/NFD/IBE+对应的儿童价格
	ChildSaleType int64 `json:"child_sale_type,omitempty" xml:"child_sale_type,omitempty"`
	// 儿童票面价，单位：分，当计算方式是3时生效，票面价必须为整数，且如果大于成人票面价，此政策无效
	ChildFixedPrice int64 `json:"child_fixed_price,omitempty" xml:"child_fixed_price,omitempty"`
	// 行李有无：null-以平台行李额为准，0-以平台行李额为准，1-无免费托运行李额，2-有免费托运行李额；
	Baggage int64 `json:"baggage,omitempty" xml:"baggage,omitempty"`
	// 去程票面价  单位：分
	OutPrice int64 `json:"out_price,omitempty" xml:"out_price,omitempty"`
	// 回程票面价	单位：分
	BackPrice int64 `json:"back_price,omitempty" xml:"back_price,omitempty"`
}

var poolPriceDto = sync.Pool{
	New: func() any {
		return new(PriceDto)
	},
}

// GetPriceDto() 从对象池中获取PriceDto
func GetPriceDto() *PriceDto {
	return poolPriceDto.Get().(*PriceDto)
}

// ReleasePriceDto 释放PriceDto
func ReleasePriceDto(v *PriceDto) {
	v.FlightPriceValues = v.FlightPriceValues[:0]
	v.Commission = 0
	v.ReturnPrice = 0
	v.MinPriceLimit = 0
	v.MaxPriceLimit = 0
	v.BiddFee = 0
	v.BiddFeePercent = 0
	v.CarryRule = 0
	v.LowestPrice = 0
	v.CalFareMethod = 0
	v.BidFee = 0
	v.BidMethod = 0
	v.ChildSaleType = 0
	v.ChildFixedPrice = 0
	v.Baggage = 0
	v.OutPrice = 0
	v.BackPrice = 0
	poolPriceDto.Put(v)
}
