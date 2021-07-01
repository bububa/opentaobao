package tvupadmin

// DeviceAdapterDto 结构体
type DeviceAdapterDto struct {
	// brandId
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// brandName
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// highestSystemVersion
	HighestSystemVersion string `json:"highest_system_version,omitempty" xml:"highest_system_version,omitempty"`
	// minimumSystemVersion
	MinimumSystemVersion string `json:"minimum_system_version,omitempty" xml:"minimum_system_version,omitempty"`
	// modelId
	ModelId int64 `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// modelName
	ModelName string `json:"model_name,omitempty" xml:"model_name,omitempty"`
	// realTypeId
	RealTypeId int64 `json:"real_type_id,omitempty" xml:"real_type_id,omitempty"`
	// realTypeName
	RealTypeName string `json:"real_type_name,omitempty" xml:"real_type_name,omitempty"`
}
