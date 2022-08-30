package crm

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
