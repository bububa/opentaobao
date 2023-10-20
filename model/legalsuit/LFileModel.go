package legalsuit

// LfileModel 结构体
type LfileModel struct {
	// 附件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 附件key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}
