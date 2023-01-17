package tbitem

// PaimaiInfo 结构体
type PaimaiInfo struct {
	// 增价幅度或降价幅度
	Increment string `json:"increment,omitempty" xml:"increment,omitempty"`
	// 降价拍的保留价
	Reserve string `json:"reserve,omitempty" xml:"reserve,omitempty"`
	// 拍品开始时间
	Start string `json:"start,omitempty" xml:"start,omitempty"`
	// 拍品结束时间
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 用户自定义的固定保证金。如果用户未传则说明用户选择使用淘宝默认的保证金模式10%，此时deposit等于0.
	Deposit int64 `json:"deposit,omitempty" xml:"deposit,omitempty"`
	// 降价拍中的降价间隔
	Interval int64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// 拍卖类型，目前包括增加拍，荷兰拍和降价拍。
	Mode int64 `json:"mode,omitempty" xml:"mode,omitempty"`
	// 对于拍卖来说可以设定有效期，这里是有效期的小时数。
	ValidHour int64 `json:"valid_hour,omitempty" xml:"valid_hour,omitempty"`
	// 对于拍卖来说可以设定有效期，这里是有效期的分钟数。
	ValidMinute int64 `json:"valid_minute,omitempty" xml:"valid_minute,omitempty"`
	// 重复上拍总次数，如果不是重复上拍的，则为0
	Repeat int64 `json:"repeat,omitempty" xml:"repeat,omitempty"`
	// 拍品起拍价
	StartPrice int64 `json:"start_price,omitempty" xml:"start_price,omitempty"`
	// 拍品封顶价（分）
	CeilPrice int64 `json:"ceil_price,omitempty" xml:"ceil_price,omitempty"`
	// 增加拍延迟周期（分钟）
	DelayInMinute int64 `json:"delay_in_minute,omitempty" xml:"delay_in_minute,omitempty"`
	// 拍卖周期（秒钟），非重复上架为（开始时间-结束时间），当为重复上拍时为一次重复上架的时间
	Period int64 `json:"period,omitempty" xml:"period,omitempty"`
	// 降价时间周期（分钟）
	Frequency int64 `json:"frequency,omitempty" xml:"frequency,omitempty"`
}
