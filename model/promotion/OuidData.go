package promotion

import (
	"sync"
)

// OuidData 结构体
type OuidData struct {
	// ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
}

var poolOuidData = sync.Pool{
	New: func() any {
		return new(OuidData)
	},
}

// GetOuidData() 从对象池中获取OuidData
func GetOuidData() *OuidData {
	return poolOuidData.Get().(*OuidData)
}

// ReleaseOuidData 释放OuidData
func ReleaseOuidData(v *OuidData) {
	v.Ouid = ""
	poolOuidData.Put(v)
}
