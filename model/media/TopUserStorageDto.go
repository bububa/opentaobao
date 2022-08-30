package media

// TopUserStorageDto 结构体
type TopUserStorageDto struct {
	// 已使用的容量，单位字节
	UsedQuota int64 `json:"used_quota,omitempty" xml:"used_quota,omitempty"`
	// 总容量，单位字节
	TotalQuota int64 `json:"total_quota,omitempty" xml:"total_quota,omitempty"`
}
