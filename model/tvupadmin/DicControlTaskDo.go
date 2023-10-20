package tvupadmin

import (
	"sync"
)

// DicControlTaskDo 结构体
type DicControlTaskDo struct {
	// 操作者
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 任务描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 任务名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 设备型号
	Devices string `json:"devices,omitempty" xml:"devices,omitempty"`
	// 任务类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// apk id
	ApkId int64 `json:"apk_id,omitempty" xml:"apk_id,omitempty"`
	// 牌照方
	License int64 `json:"license,omitempty" xml:"license,omitempty"`
}

var poolDicControlTaskDo = sync.Pool{
	New: func() any {
		return new(DicControlTaskDo)
	},
}

// GetDicControlTaskDo() 从对象池中获取DicControlTaskDo
func GetDicControlTaskDo() *DicControlTaskDo {
	return poolDicControlTaskDo.Get().(*DicControlTaskDo)
}

// ReleaseDicControlTaskDo 释放DicControlTaskDo
func ReleaseDicControlTaskDo(v *DicControlTaskDo) {
	v.Operator = ""
	v.Description = ""
	v.Name = ""
	v.Uuid = ""
	v.Devices = ""
	v.Type = 0
	v.ApkId = 0
	v.License = 0
	poolDicControlTaskDo.Put(v)
}
