package ieagency

// IeBookPnrVo 结构体
type IeBookPnrVo struct {
	// pnr id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// pnr类型
	PnrType string `json:"pnr_type,omitempty" xml:"pnr_type,omitempty"`
	// pnr值
	PnrNo string `json:"pnr_no,omitempty" xml:"pnr_no,omitempty"`
}
