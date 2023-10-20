package westcrm

// AlibabaWestcrmGradeGetData 结构体
type AlibabaWestcrmGradeGetData struct {
	// 等级名称
	GradeName string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
	// 等级列表
	GradeNum int64 `json:"grade_num,omitempty" xml:"grade_num,omitempty"`
	// 等级id
	GradeId int64 `json:"grade_id,omitempty" xml:"grade_id,omitempty"`
}
