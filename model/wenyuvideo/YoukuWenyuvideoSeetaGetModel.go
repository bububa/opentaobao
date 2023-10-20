package wenyuvideo

// YoukuwenyuvideoseetagetModel 结构体
type YoukuwenyuvideoseetagetModel struct {
	// logo数组
	Logos []string `json:"logos,omitempty" xml:"logos>string,omitempty"`
	// 名字数组
	Names []string `json:"names,omitempty" xml:"names>string,omitempty"`
	// 片段数组
	Segments []Segments `json:"segments,omitempty" xml:"segments>segments,omitempty"`
}
