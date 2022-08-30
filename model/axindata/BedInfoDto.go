package axindata

// BedInfoDto 结构体
type BedInfoDto struct {
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 床型0-大床1-双床2-单人床
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 床宽
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 床长
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 数量
	BedNum int64 `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
}
