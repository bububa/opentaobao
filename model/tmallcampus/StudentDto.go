package tmallcampus

// StudentDto 结构体
type StudentDto struct {
	// 学校名称
	Campus string `json:"campus,omitempty" xml:"campus,omitempty"`
	// 认证状态(0:从未参与过认证的| 1: 算法预判通过 | 2:算法预判失败 | 3:人工审核中 | 4:人工审核通过 | 5:人工审核失败)
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
