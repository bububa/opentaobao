package crm

import (
	"sync"
)

// HsMemberInfoDto 结构体
type HsMemberInfoDto struct {
	// 版本拓展信息
	SnapshotInfo string `json:"snapshot_info,omitempty" xml:"snapshot_info,omitempty"`
	// 等级名称
	GradeName string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
	// ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// 记录最后修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 等级编码
	Grade int64 `json:"grade,omitempty" xml:"grade,omitempty"`
}

var poolHsMemberInfoDto = sync.Pool{
	New: func() any {
		return new(HsMemberInfoDto)
	},
}

// GetHsMemberInfoDto() 从对象池中获取HsMemberInfoDto
func GetHsMemberInfoDto() *HsMemberInfoDto {
	return poolHsMemberInfoDto.Get().(*HsMemberInfoDto)
}

// ReleaseHsMemberInfoDto 释放HsMemberInfoDto
func ReleaseHsMemberInfoDto(v *HsMemberInfoDto) {
	v.SnapshotInfo = ""
	v.GradeName = ""
	v.Ouid = ""
	v.GmtModified = ""
	v.Grade = 0
	poolHsMemberInfoDto.Put(v)
}
