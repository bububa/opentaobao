package omniorder

// StoreDeliverConfig 结构体
type StoreDeliverConfig struct {
	// 当activity为true时返回,活动开始时间
	ActivityStartTime string `json:"activity_start_time,omitempty" xml:"activity_start_time,omitempty"`
	// 当activity为true时返回，活动结束时间
	ActivityEndTime string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
	// 接单时间段，格式为 &#34;09:00-12:00&#34;, &#34;&#34; 表示一直开启
	WorkingTime string `json:"working_time,omitempty" xml:"working_time,omitempty"`
	// 派单时间，格式为：[{&#34;startTime&#34;:&#34;08:40&#34;,&#34;endTime&#34;:&#34;12:20&#34;},{&#34;startTime&#34;:&#34;18:00&#34;,&#34;endTime&#34;:&#34;22:00&#34;}]
	DispatchTimeRange string `json:"dispatch_time_range,omitempty" xml:"dispatch_time_range,omitempty"`
	// 优先级，取值0-10，0优先级最大，10优先级最小
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 每日接单阈值
	DeliverThreshold int64 `json:"deliver_threshold,omitempty" xml:"deliver_threshold,omitempty"`
	// 是否是活动期
	Activity bool `json:"activity,omitempty" xml:"activity,omitempty"`
}
