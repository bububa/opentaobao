package tvupadmin

// AppDto 结构体
type AppDto struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// packageName
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 应用ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
