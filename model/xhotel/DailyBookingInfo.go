package xhotel

// DailyBookingInfo 结构体
type DailyBookingInfo struct {
	// 一个星期内有效性约束。1-7 对应周一到周日，传入的值比如[1,6]，就表示星期一和星期六营销生效
	ValidWeeks []string `json:"valid_weeks,omitempty" xml:"valid_weeks>string,omitempty"`
	// 不可用日期，开始日期和结束日期: from--to  只有一天的场景，from和to传同一天； 默认：空，代表无时间限制；
	InvalidDates []InvalidDate `json:"invalid_dates,omitempty" xml:"invalid_dates>invalid_date,omitempty"`
	// 可入住的起始时间，不填默认一年，一年后自动续期
	CheckInFrom string `json:"check_in_from,omitempty" xml:"check_in_from,omitempty"`
	// 可入住的结束时间，不填默认一年，一年后自动续期
	CheckInTo string `json:"check_in_to,omitempty" xml:"check_in_to,omitempty"`
	// 折扣比例，填30就意味着原价的30%，也就是打3折。数字范围限定在10-95之间
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
}
