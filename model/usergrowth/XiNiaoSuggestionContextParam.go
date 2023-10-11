package usergrowth

// XiNiaoSuggestionContextParam 结构体
type XiNiaoSuggestionContextParam struct {
	// open_id
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// open_id
	OpenId2 string `json:"open_id2,omitempty" xml:"open_id2,omitempty"`
	// 站点id
	UnionCode string `json:"union_code,omitempty" xml:"union_code,omitempty"`
	// 包裹数量
	PackageNum int64 `json:"package_num,omitempty" xml:"package_num,omitempty"`
}
