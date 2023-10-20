package dt

import (
	"sync"
)

// OdpsImportConfig 结构体
type OdpsImportConfig struct {
	// odps表名
	OdpsTable string `json:"odps_table,omitempty" xml:"odps_table,omitempty"`
	// odps账号
	OdpsAccount string `json:"odps_account,omitempty" xml:"odps_account,omitempty"`
	// odps项目
	OdpsProject string `json:"odps_project,omitempty" xml:"odps_project,omitempty"`
	// odps表标记
	OdpsTablePartition string `json:"odps_table_partition,omitempty" xml:"odps_table_partition,omitempty"`
}

var poolOdpsImportConfig = sync.Pool{
	New: func() any {
		return new(OdpsImportConfig)
	},
}

// GetOdpsImportConfig() 从对象池中获取OdpsImportConfig
func GetOdpsImportConfig() *OdpsImportConfig {
	return poolOdpsImportConfig.Get().(*OdpsImportConfig)
}

// ReleaseOdpsImportConfig 释放OdpsImportConfig
func ReleaseOdpsImportConfig(v *OdpsImportConfig) {
	v.OdpsTable = ""
	v.OdpsAccount = ""
	v.OdpsProject = ""
	v.OdpsTablePartition = ""
	poolOdpsImportConfig.Put(v)
}
