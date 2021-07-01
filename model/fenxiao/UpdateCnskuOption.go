package fenxiao

// UpdateCnskuOption 结构体
type UpdateCnskuOption struct {
	// 是否同步到wms, 为空时默认下发
	SyncWms bool `json:"sync_wms,omitempty" xml:"sync_wms,omitempty"`
}
