package seaking

// AlibabaseakingservicepackResult 结构体
type AlibabaseakingservicepackResult struct {
	// 到期时间
	ValidateTo string `json:"validate_to,omitempty" xml:"validate_to,omitempty"`
	// 权限包名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 权限包id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
