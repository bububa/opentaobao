package tmallcarenter

// VersionVehicleInfoOriginalDto 结构体
type VersionVehicleInfoOriginalDto struct {
	// 换代
	Replacement string `json:"replacement,omitempty" xml:"replacement,omitempty"`
	// 销售名称
	SalesName string `json:"sales_name,omitempty" xml:"sales_name,omitempty"`
	// 产地
	Origin string `json:"origin,omitempty" xml:"origin,omitempty"`
	// 品牌
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 后轮毂规格
	RearTyre string `json:"rear_tyre,omitempty" xml:"rear_tyre,omitempty"`
	// 最大功率
	MaxPower string `json:"max_power,omitempty" xml:"max_power,omitempty"`
	// 生产厂家
	ManufactureName string `json:"manufacture_name,omitempty" xml:"manufacture_name,omitempty"`
	// 版本压缩ID
	VersionCid string `json:"version_cid,omitempty" xml:"version_cid,omitempty"`
	// 年款
	SalesYear string `json:"sales_year,omitempty" xml:"sales_year,omitempty"`
	// 发动机型号
	EngineNo string `json:"engine_no,omitempty" xml:"engine_no,omitempty"`
	// 车型
	ModelName string `json:"model_name,omitempty" xml:"model_name,omitempty"`
	// 前轮毂规格
	FrontTyre string `json:"front_tyre,omitempty" xml:"front_tyre,omitempty"`
	// 排量
	Displacement string `json:"displacement,omitempty" xml:"displacement,omitempty"`
	// 驱动方式
	DriveModel string `json:"drive_model,omitempty" xml:"drive_model,omitempty"`
	// 底盘号
	ChassisNum string `json:"chassis_num,omitempty" xml:"chassis_num,omitempty"`
	// 车身类型
	BodyModel string `json:"body_model,omitempty" xml:"body_model,omitempty"`
	// 车系
	LineName string `json:"line_name,omitempty" xml:"line_name,omitempty"`
	// 车辆类型
	VehicleType string `json:"vehicle_type,omitempty" xml:"vehicle_type,omitempty"`
	// 停产年份
	EndYear int64 `json:"end_year,omitempty" xml:"end_year,omitempty"`
	// 生产年份
	ProductiveYear int64 `json:"productive_year,omitempty" xml:"productive_year,omitempty"`
}
