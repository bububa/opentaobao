package ascp

// DeliveryResourceCreateRequest 结构体
type DeliveryResourceCreateRequest struct {
	// 幂等ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 配品牌名称（资源名称）
	DeliveryName string `json:"delivery_name,omitempty" xml:"delivery_name,omitempty"`
	// 配品牌
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 服务商配资源唯一编码
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
	// 平台快递资源
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 联系人姓名
	ConName string `json:"con_name,omitempty" xml:"con_name,omitempty"`
	// 联系人电话
	ConPhone string `json:"con_phone,omitempty" xml:"con_phone,omitempty"`
	// 时间戳(毫秒级)
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 配资源类型，枚举： 20：快递 10：即时配
	DeliveryType int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
}
