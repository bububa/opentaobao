package foodscan

// FeetModelMobileReportVo 
type FeetModelMobileReportVo struct {
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 扫描ID
    ScanId   string `json:"scan_id,omitempty" xml:"scan_id,omitempty"`
    // 脚长
    FootLength   string `json:"foot_length,omitempty" xml:"foot_length,omitempty"`
    // 脚宽
    FootWidth   string `json:"foot_width,omitempty" xml:"foot_width,omitempty"`
    // 性别  1：男 2：女
    Gender   int64 `json:"gender,omitempty" xml:"gender,omitempty"`
}
