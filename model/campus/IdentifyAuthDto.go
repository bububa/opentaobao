package campus

// IdentifyAuthDto 结构体
type IdentifyAuthDto struct {
	// []
	VoucherList []VoucherDto `json:"voucher_list,omitempty" xml:"voucher_list>voucher_dto,omitempty"`
	// 指定鉴权类型
	AuthTypeEnum string `json:"auth_type_enum,omitempty" xml:"auth_type_enum,omitempty"`
	// 属性json
	PropertiesJson string `json:"properties_json,omitempty" xml:"properties_json,omitempty"`
	// 应用code
	AppCode string `json:"app_code,omitempty" xml:"app_code,omitempty"`
	// 子设备ID
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 签名
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 时间戳
	TimeStamp int64 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
}
