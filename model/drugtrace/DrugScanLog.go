package drugtrace

// DrugScanLog 结构体
type DrugScanLog struct {
	// 包装规格
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 溯源码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 有效期
	ExpiryDate string `json:"expiry_date,omitempty" xml:"expiry_date,omitempty"`
}
