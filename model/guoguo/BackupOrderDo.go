package guoguo

import (
	"sync"
)

// BackupOrderDo 结构体
type BackupOrderDo struct {
	// 返回成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBackupOrderDo = sync.Pool{
	New: func() any {
		return new(BackupOrderDo)
	},
}

// GetBackupOrderDo() 从对象池中获取BackupOrderDo
func GetBackupOrderDo() *BackupOrderDo {
	return poolBackupOrderDo.Get().(*BackupOrderDo)
}

// ReleaseBackupOrderDo 释放BackupOrderDo
func ReleaseBackupOrderDo(v *BackupOrderDo) {
	v.Success = false
	poolBackupOrderDo.Put(v)
}
