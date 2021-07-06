package tmallcarenter

// ChasisVehicleInfoOriginalDto 结构体
type ChasisVehicleInfoOriginalDto struct {
	// 换代
	Replacement string `json:"replacement,omitempty" xml:"replacement,omitempty"`
	// 产地
	Origin string `json:"origin,omitempty" xml:"origin,omitempty"`
	// 品牌
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 最大功率
	MaxPower string `json:"max_power,omitempty" xml:"max_power,omitempty"`
	// 厂家
	ManufactureName string `json:"manufacture_name,omitempty" xml:"manufacture_name,omitempty"`
	// 发动机型号
	EngineNo string `json:"engine_no,omitempty" xml:"engine_no,omitempty"`
	// 车型
	ModelName string `json:"model_name,omitempty" xml:"model_name,omitempty"`
	// 燃料类型
	FuelType string `json:"fuel_type,omitempty" xml:"fuel_type,omitempty"`
	// 排量
	Displacement string `json:"displacement,omitempty" xml:"displacement,omitempty"`
	// 驱动方式
	DriveModel string `json:"drive_model,omitempty" xml:"drive_model,omitempty"`
	// 底盘ID
	ChasisCid string `json:"chasis_cid,omitempty" xml:"chasis_cid,omitempty"`
	// 底盘号
	ChassisNum string `json:"chassis_num,omitempty" xml:"chassis_num,omitempty"`
	// 车系
	LineName string `json:"line_name,omitempty" xml:"line_name,omitempty"`
	// 停产年份
	EndYear int64 `json:"end_year,omitempty" xml:"end_year,omitempty"`
	// 生产年份
	ProductiveYear int64 `json:"productive_year,omitempty" xml:"productive_year,omitempty"`
}
