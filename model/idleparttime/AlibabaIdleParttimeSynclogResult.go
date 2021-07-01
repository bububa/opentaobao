package idleparttime

// AlibabaIdleParttimeSynclogResult 结构体
type AlibabaIdleParttimeSynclogResult struct {
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回数据结构
	Data *AlibabaIdleParttimeSynclogData `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
