package tbk

// TopApiAfOrderOption 结构体
type TopApiAfOrderOption struct {
	// 查询开始时间，以taoke订单创建时间开始
	Starttime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 渠道关系id
	Relationid int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 子订单号
	Tbtradeid int64 `json:"tb_trade_id,omitempty" xml:"tb_trade_id,omitempty"`
	// 此参数不再使用，请勿入参
	Tbtradeparentid int64 `json:"tb_trade_parent_id,omitempty" xml:"tb_trade_parent_id,omitempty"`
	// pagesize
	Pagesize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// pageNo
	Pageno int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 查询时间跨度，不超过30天，单位是天
	Span int64 `json:"span,omitempty" xml:"span,omitempty"`
	// 此参数不再使用，请勿入参
	Specialid int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
	// pid中的第三段，adzoneId
	Adzoneid int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// pid中的第二段，siteId
	Siteid int64 `json:"site_id,omitempty" xml:"site_id,omitempty"`
	// 此参数不再使用，请勿入参
	Violationtype int64 `json:"violation_type,omitempty" xml:"violation_type,omitempty"`
	// 此参数不再使用，请勿入参
	Punishstatus int64 `json:"punish_status,omitempty" xml:"punish_status,omitempty"`
}
