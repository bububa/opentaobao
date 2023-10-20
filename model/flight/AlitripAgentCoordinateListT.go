package flight

// AlitripagentcoordinatelistT 结构体
type AlitripagentcoordinatelistT struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 跟进人名称
	CurrentFollowerName string `json:"current_follower_name,omitempty" xml:"current_follower_name,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 协同履约截止时间
	ServeDeadline string `json:"serve_deadline,omitempty" xml:"serve_deadline,omitempty"`
	// 飞猪订单号
	CorrelationOutOrderId string `json:"correlation_out_order_id,omitempty" xml:"correlation_out_order_id,omitempty"`
	// 协同单ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 协同单问题类型：1, &#34;加订婴童&#34;； 2, &#34;加订婴童待出票&#34;； 3, &#34;自营履约系统转人工&#34;； 4, &#34;正向履约返现转人工&#34;； 5, &#34;催出票&#34;； 6, &#34;催退票&#34;； 7, &#34;催改签回填&#34;； 8, &#34;拦截出票&#34;； 9, &#34;系统主动二次回填&#34;； 10, &#34;商家主动二次回填&#34;； 13, &#34;追款&#34;； 11, &#34;改名/改证件询价&#34;； 12, &#34;改名/改证件&#34;； 14, &#34;补退&#34;； 15, &#34;宠物托运&#34;； 16, &#34;特殊餐食&#34;； 17, &#34;轮椅/担架&#34;； 18, &#34;婴儿摇篮&#34;； 19, &#34;加订行李&#34;； 20, &#34;航变清Q转人工&#34;； 22,&#34;无法出票&#34;； 23,&#34;线下编码预定&#34;； 24,&#34;申请关闭验真&#34;； 25, &#34;拦截改签&#34;； 26, &#34;催改签出票&#34;；
	CaseType int64 `json:"case_type,omitempty" xml:"case_type,omitempty"`
	// 协同单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否催促过，0:未催促，1:已催促
	Urge int64 `json:"urge,omitempty" xml:"urge,omitempty"`
	// 1:出票,2:退票,3:改签,4:航变
	CorrelationBizType int64 `json:"correlation_biz_type,omitempty" xml:"correlation_biz_type,omitempty"`
}
