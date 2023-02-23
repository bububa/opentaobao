package train

// ChangeApplyOrderDto 结构体
type ChangeApplyOrderDto struct {
	// 原票到达站三字码
	OriginalToStationCode string `json:"original_to_station_code,omitempty" xml:"original_to_station_code,omitempty"`
	// 改签票出发站名称
	ChangeFromStation string `json:"change_from_station,omitempty" xml:"change_from_station,omitempty"`
	// 原票123N6订单号
	OriginalSequenceNo string `json:"original_sequence_no,omitempty" xml:"original_sequence_no,omitempty"`
	// 改签票出发站三字码
	ChangeFromStationCode string `json:"change_from_station_code,omitempty" xml:"change_from_station_code,omitempty"`
	// 改签票到达站名称
	ChangeToStation string `json:"change_to_station,omitempty" xml:"change_to_station,omitempty"`
	// 改签出票超时时间
	ChangeIssueTimeout string `json:"change_issue_timeout,omitempty" xml:"change_issue_timeout,omitempty"`
	// 原票发车时间
	OriginalFromDateTime string `json:"original_from_date_time,omitempty" xml:"original_from_date_time,omitempty"`
	// 改签票达到时间
	ChangeToDateTime string `json:"change_to_date_time,omitempty" xml:"change_to_date_time,omitempty"`
	// 原票车次号
	OriginalTrainCode string `json:"original_train_code,omitempty" xml:"original_train_code,omitempty"`
	// 原票发车日期
	OriginalTrainDate string `json:"original_train_date,omitempty" xml:"original_train_date,omitempty"`
	// 原票出发站名称
	OriginalFromStation string `json:"original_from_station,omitempty" xml:"original_from_station,omitempty"`
	// 改签票达到站三字码
	ChangeToStationCode string `json:"change_to_station_code,omitempty" xml:"change_to_station_code,omitempty"`
	// 原票达到站名称
	OriginalToStation string `json:"original_to_station,omitempty" xml:"original_to_station,omitempty"`
	// 原票出发站三字码
	OriginalFromStationCode string `json:"original_from_station_code,omitempty" xml:"original_from_station_code,omitempty"`
	// 原票到达时间
	OriginalToDateTime string `json:"original_to_date_time,omitempty" xml:"original_to_date_time,omitempty"`
	// 改签票发车时间
	ChangeFromDateTime string `json:"change_from_date_time,omitempty" xml:"change_from_date_time,omitempty"`
	// 改签票发车日期
	ChangeTrainDate string `json:"change_train_date,omitempty" xml:"change_train_date,omitempty"`
	// 改签票车次号
	ChangeTrainCode string `json:"change_train_code,omitempty" xml:"change_train_code,omitempty"`
	// 淘宝主单单号
	TpId int64 `json:"tp_id,omitempty" xml:"tp_id,omitempty"`
	// 代理商id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 改签票类型 0:改签 1:变更到站
	ChangeTicketType int64 `json:"change_ticket_type,omitempty" xml:"change_ticket_type,omitempty"`
	// 改签申请单id
	ChangeApplyId int64 `json:"change_apply_id,omitempty" xml:"change_apply_id,omitempty"`
	// 现金结算标示 0:现金结算 1:电子结算
	SettlementMode int64 `json:"settlement_mode,omitempty" xml:"settlement_mode,omitempty"`
	// 改签类型 0:低改高 1:平改 2:高改低
	ChangeType int64 `json:"change_type,omitempty" xml:"change_type,omitempty"`
	// 整单票价总金额
	TicketPriceAll int64 `json:"ticket_price_all,omitempty" xml:"ticket_price_all,omitempty"`
	// ttp单号
	TtpId int64 `json:"ttp_id,omitempty" xml:"ttp_id,omitempty"`
	// 票数量
	TicketTotalNum int64 `json:"ticket_total_num,omitempty" xml:"ticket_total_num,omitempty"`
	// 改签申请单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
