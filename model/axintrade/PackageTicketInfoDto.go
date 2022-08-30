package axintrade

// PackageTicketInfoDto 结构体
type PackageTicketInfoDto struct {
	// 门票产品信息
	TicketInfo *TicketInfoDto `json:"ticket_info,omitempty" xml:"ticket_info,omitempty"`
	// 门票张数
	TicketNum int64 `json:"ticket_num,omitempty" xml:"ticket_num,omitempty"`
}
