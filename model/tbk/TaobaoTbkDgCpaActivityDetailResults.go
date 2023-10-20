package tbk

// TaobaotbkdgcpaactivitydetailResults 结构体
type TaobaotbkdgcpaactivitydetailResults struct {
	// 奖励明细数据，KV结构。字段释义见文档：https://www.yuque.com/docs/share/7ecf8cf1-7f99-4633-a2ed-f9b6f8116af5?#
	Fielddetail string `json:"field_detail,omitempty" xml:"field_detail,omitempty"`
	// 明细类型，1：预估明细，2：结算明细
	Calctype int64 `json:"calc_type,omitempty" xml:"calc_type,omitempty"`
	// 明细记录主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
