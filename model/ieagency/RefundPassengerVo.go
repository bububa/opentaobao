package ieagency

// RefundPassengerVo 结构体
type RefundPassengerVo struct {
	// 乘机人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 乘机人ID
	PassenerId int64 `json:"passener_id,omitempty" xml:"passener_id,omitempty"`
	// 乘机人类型(Adult(0, &#34;成人&#34;),     Child(1, &#34;儿童&#34;),     StudentAbroad(2, &#34;留学生&#34;),     Infant(3, &#34;婴儿&#34;)
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
}
