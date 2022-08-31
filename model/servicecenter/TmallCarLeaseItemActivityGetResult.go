package servicecenter

// TmallCarLeaseItemActivityGetResult 结构体
type TmallCarLeaseItemActivityGetResult struct {
	// 动态返回扩展参数：&lt;br/&gt;extInfo:扩展参数字符串
	Piggyback string `json:"piggyback,omitempty" xml:"piggyback,omitempty"`
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 活动返回列表对象&lt;br/&gt; itemId：商品ID &lt;br/&gt; preEndTime：预热结束时间&lt;br/&gt; preStartTime：预热开始时间&lt;br/&gt; &lt;br/&gt; extendInfo：扩展信息&lt;br/&gt; refActivityId:外部活动ID&lt;br/&gt; &lt;br/&gt; timeRangeList：活动时间范围对象&lt;br/&gt; startTime：活动开始时间&lt;br/&gt; endTime：活动结束时间&lt;br/&gt; amount：名额&lt;br/&gt; extendInfo：扩展信息&lt;br/&gt;
	Object string `json:"object,omitempty" xml:"object,omitempty"`
	// 当前时间
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// 耗费时间
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
