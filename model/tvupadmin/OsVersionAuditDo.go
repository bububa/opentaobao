package tvupadmin

import (
	"sync"
)

// OsVersionAuditDo 结构体
type OsVersionAuditDo struct {
	// 升级包列表
	OsRomList []OsRomDo `json:"os_rom_list,omitempty" xml:"os_rom_list>os_rom_do,omitempty"`
	// 型号名
	ModelName string `json:"model_name,omitempty" xml:"model_name,omitempty"`
	// 内部型号名
	RealTypeName string `json:"real_type_name,omitempty" xml:"real_type_name,omitempty"`
	// 版本号
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 发布说明
	ReleaseNote string `json:"release_note,omitempty" xml:"release_note,omitempty"`
	// 审核状态
	AuditStatus string `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 审核时间
	GmtAudit string `json:"gmt_audit,omitempty" xml:"gmt_audit,omitempty"`
	// 主键ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolOsVersionAuditDo = sync.Pool{
	New: func() any {
		return new(OsVersionAuditDo)
	},
}

// GetOsVersionAuditDo() 从对象池中获取OsVersionAuditDo
func GetOsVersionAuditDo() *OsVersionAuditDo {
	return poolOsVersionAuditDo.Get().(*OsVersionAuditDo)
}

// ReleaseOsVersionAuditDo 释放OsVersionAuditDo
func ReleaseOsVersionAuditDo(v *OsVersionAuditDo) {
	v.OsRomList = v.OsRomList[:0]
	v.ModelName = ""
	v.RealTypeName = ""
	v.Version = ""
	v.ReleaseNote = ""
	v.AuditStatus = ""
	v.StatusDesc = ""
	v.GmtCreate = ""
	v.GmtModify = ""
	v.GmtAudit = ""
	v.Id = 0
	poolOsVersionAuditDo.Put(v)
}
