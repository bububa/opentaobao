package jipiao

// BbSyncOrderDto 结构体
type BbSyncOrderDto struct {
	// 申请单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 改签后舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 改签前舱位
	LastCabin string `json:"last_cabin,omitempty" xml:"last_cabin,omitempty"`
	// 改签备注信息
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 改签后航班信息
	ModifyAfterFlight *Flight `json:"modify_after_flight,omitempty" xml:"modify_after_flight,omitempty"`
	// 改签前航班信息
	ModifyBeforeFlight *Flight `json:"modify_before_flight,omitempty" xml:"modify_before_flight,omitempty"`
	// 乘机人信息
	Passenger *Passenger `json:"passenger,omitempty" xml:"passenger,omitempty"`
	// 改签前PNR编码
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 改签前票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}
