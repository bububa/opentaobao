package fenxiao

// AddCnskuOption 
type AddCnskuOption struct {

    // 是否同步到wms, 为空时默认下发
    SyncWms   bool `json:"sync_wms,omitempty"`

}
