package simba

// AdgroupOcpcVo 结构体
type AdgroupOcpcVo struct {
	// OCPC溢价比例
	OcpcRatio int64 `json:"ocpc_ratio,omitempty" xml:"ocpc_ratio,omitempty"`
	// OCPC是否开启，false:否，true:是
	EnableOcpc bool `json:"enable_ocpc,omitempty" xml:"enable_ocpc,omitempty"`
}
