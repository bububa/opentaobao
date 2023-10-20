package alihealth2

import (
	"sync"
)

// CommonRequest4Top 结构体
type CommonRequest4Top struct {
	// hosNo:渠道医院ID,deptNo:渠道科室ID,deptName:科室名称,status:状态
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
}

var poolCommonRequest4Top = sync.Pool{
	New: func() any {
		return new(CommonRequest4Top)
	},
}

// GetCommonRequest4Top() 从对象池中获取CommonRequest4Top
func GetCommonRequest4Top() *CommonRequest4Top {
	return poolCommonRequest4Top.Get().(*CommonRequest4Top)
}

// ReleaseCommonRequest4Top 释放CommonRequest4Top
func ReleaseCommonRequest4Top(v *CommonRequest4Top) {
	v.BizContent = ""
	poolCommonRequest4Top.Put(v)
}
