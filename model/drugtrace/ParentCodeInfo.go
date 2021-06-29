package drugtrace

// ParentCodeInfo 
type ParentCodeInfo struct {
    // 码级别
    PackageLevel   int64 `json:"package_level,omitempty" xml:"package_level,omitempty"`
    // 父码
    ParentCode   string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
}
