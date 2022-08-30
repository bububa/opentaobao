package wdk

// SameTownBox 结构体
type SameTownBox struct {
	// 同城包裹列表
	SameTownPackages []Sametownpackages `json:"same_town_packages,omitempty" xml:"same_town_packages>sametownpackages,omitempty"`
	// 是否测试 1:测试 0:非测试
	IsTest string `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 扩展信息  json格式
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 箱号
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// 箱类型 NORMAL:周转箱 COLD:冷链箱 NONE:原箱
	ContainerType string `json:"container_type,omitempty" xml:"container_type,omitempty"`
	// 箱号
	MaterialCode string `json:"material_code,omitempty" xml:"material_code,omitempty"`
}
