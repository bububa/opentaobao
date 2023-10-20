package fenxiao

import (
	"sync"
)

// FenxiaoGrade 结构体
type FenxiaoGrade struct {
	// 分销商等级名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 记录创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 记录最后修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 主键
	GradeId int64 `json:"grade_id,omitempty" xml:"grade_id,omitempty"`
}

var poolFenxiaoGrade = sync.Pool{
	New: func() any {
		return new(FenxiaoGrade)
	},
}

// GetFenxiaoGrade() 从对象池中获取FenxiaoGrade
func GetFenxiaoGrade() *FenxiaoGrade {
	return poolFenxiaoGrade.Get().(*FenxiaoGrade)
}

// ReleaseFenxiaoGrade 释放FenxiaoGrade
func ReleaseFenxiaoGrade(v *FenxiaoGrade) {
	v.Name = ""
	v.Created = ""
	v.Modified = ""
	v.GradeId = 0
	poolFenxiaoGrade.Put(v)
}
