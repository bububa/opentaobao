package xhotel

// Adsslrservicedatalist 结构体
type Adsslrservicedatalist struct {
	// 查无订单率
	NoOrdCnt1dRateDouble string `json:"no_ord_cnt1d_rate_double,omitempty" xml:"no_ord_cnt1d_rate_double,omitempty"`
	// 退款率
	PreRefundCnt1mRateDouble string `json:"pre_refund_cnt1m_rate_double,omitempty" xml:"pre_refund_cnt1m_rate_double,omitempty"`
	// 国内预订成功率
	BookingSuccRateDomesticDouble string `json:"booking_succ_rate_domestic_double,omitempty" xml:"booking_succ_rate_domestic_double,omitempty"`
	// 查询日期
	ReportDate string `json:"report_date,omitempty" xml:"report_date,omitempty"`
	// 预订成功率国际
	BookingSuccRateOverseasDouble string `json:"booking_succ_rate_overseas_double,omitempty" xml:"booking_succ_rate_overseas_double,omitempty"`
	// 到店无房率
	NoRoomCnt1dRateDouble string `json:"no_room_cnt1d_rate_double,omitempty" xml:"no_room_cnt1d_rate_double,omitempty"`
	// 卖家原因退款率
	PreSlrRefundCnt1mRateDouble string `json:"pre_slr_refund_cnt1m_rate_double,omitempty" xml:"pre_slr_refund_cnt1m_rate_double,omitempty"`
	// 查无订单数
	NoOrdCnt1d int64 `json:"no_ord_cnt1d,omitempty" xml:"no_ord_cnt1d,omitempty"`
	// 90天销量
	SalesCountIn90days int64 `json:"sales_count_in90days,omitempty" xml:"sales_count_in90days,omitempty"`
	// 国际平均响应时长
	RespDurationOverseas int64 `json:"resp_duration_overseas,omitempty" xml:"resp_duration_overseas,omitempty"`
	// 到店无房订单数
	NoRoomCnt1d int64 `json:"no_room_cnt1d,omitempty" xml:"no_room_cnt1d,omitempty"`
	// 每日订单数
	SalesCountInDay int64 `json:"sales_count_in_day,omitempty" xml:"sales_count_in_day,omitempty"`
}
