package westcrm

import (
	"sync"
)

// AlibabaWestcrmGradeGetData 结构体
type AlibabaWestcrmGradeGetData struct {
	// 等级名称
	GradeName string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
	// 等级列表
	GradeNum int64 `json:"grade_num,omitempty" xml:"grade_num,omitempty"`
	// 等级id
	GradeId int64 `json:"grade_id,omitempty" xml:"grade_id,omitempty"`
}

var poolAlibabaWestcrmGradeGetData = sync.Pool{
	New: func() any {
		return new(AlibabaWestcrmGradeGetData)
	},
}

// GetAlibabaWestcrmGradeGetData() 从对象池中获取AlibabaWestcrmGradeGetData
func GetAlibabaWestcrmGradeGetData() *AlibabaWestcrmGradeGetData {
	return poolAlibabaWestcrmGradeGetData.Get().(*AlibabaWestcrmGradeGetData)
}

// ReleaseAlibabaWestcrmGradeGetData 释放AlibabaWestcrmGradeGetData
func ReleaseAlibabaWestcrmGradeGetData(v *AlibabaWestcrmGradeGetData) {
	v.GradeName = ""
	v.GradeNum = 0
	v.GradeId = 0
	poolAlibabaWestcrmGradeGetData.Put(v)
}
