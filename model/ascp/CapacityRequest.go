package ascp

// CapacityRequest 结构体
type CapacityRequest struct {
	// 行政地址id（菜鸟地址库id）
	AddressIds []string `json:"address_ids,omitempty" xml:"address_ids>string,omitempty"`
	// 中文地址信息
	AddressNames []AddressName `json:"address_names,omitempty" xml:"address_names>address_name,omitempty"`
	// 日常产能信息
	DailyCapacityInfos []CapacityInfo `json:"daily_capacity_infos,omitempty" xml:"daily_capacity_infos>capacity_info,omitempty"`
	// 指定日期产能信息
	SpecifyCapacityInfos []SpecifyCapacityInfo `json:"specify_capacity_infos,omitempty" xml:"specify_capacity_infos>specify_capacity_info,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 配资源编码
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
	// 服务类型：1-上门取退
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 能力：1-上门取退
	AbilityType string `json:"ability_type,omitempty" xml:"ability_type,omitempty"`
	// 服务范围地址类型：1-行政地址；2-电子围栏
	ServiceScopeType string `json:"service_scope_type,omitempty" xml:"service_scope_type,omitempty"`
	// 如果传入为行政地址，行政地址传入类型 1- 菜鸟地址库ID 传入 2- 中文地址传入
	AddressType string `json:"address_type,omitempty" xml:"address_type,omitempty"`
	// 电子围栏外部ID64	条件必填（service_scope_type=2时，必填）
	RegionCode string `json:"region_code,omitempty" xml:"region_code,omitempty"`
	// 时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 产能更新方式：1-全量更新；2-部分更新 ● 电子围栏产能首次同步时，需选择全量更新 ● 选择全量更新时，日常及指定日期时间段产能均做全量覆盖更新； ● 选择部分更新时，日常及指定日期时间段产能仅对传入的时间段的产能进行更新，其他时间段不做处理，维持原状
	UpdateMethod int64 `json:"update_method,omitempty" xml:"update_method,omitempty"`
}
