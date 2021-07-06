package tvupadmin

// Deviceadapterlist 结构体
type Deviceadapterlist struct {
	// 设备最小版本号
	MinimumSystemVersion string `json:"minimum_system_version,omitempty" xml:"minimum_system_version,omitempty"`
	// 设备最大版本号
	HighestSystemVersion string `json:"highest_system_version,omitempty" xml:"highest_system_version,omitempty"`
	// 设备真实名称
	RealTypeName string `json:"real_type_name,omitempty" xml:"real_type_name,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 设备名称
	ModelName string `json:"model_name,omitempty" xml:"model_name,omitempty"`
	// 适配设备真实类型ID
	RealTypeId int64 `json:"real_type_id,omitempty" xml:"real_type_id,omitempty"`
	// 品牌ID
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 设备ID
	ModelId int64 `json:"model_id,omitempty" xml:"model_id,omitempty"`
}
