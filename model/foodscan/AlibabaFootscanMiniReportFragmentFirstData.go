package foodscan

// AlibabaFootscanMiniReportFragmentFirstData 结构体
type AlibabaFootscanMiniReportFragmentFirstData struct {
	// 扫描ID
	ScanId string `json:"scan_id,omitempty" xml:"scan_id,omitempty"`
	// 测量结果为0
	IsZero bool `json:"is_zero,omitempty" xml:"is_zero,omitempty"`
}
