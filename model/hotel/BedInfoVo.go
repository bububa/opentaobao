package hotel

// BedInfoVo 结构体
type BedInfoVo struct {
	// 床型
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 床长
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 长款
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 床数量
	BedNum int64 `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
}
