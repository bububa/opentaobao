package vaccin

// SubmitVcRegisterRequest 结构体
type SubmitVcRegisterRequest struct {
	// cdc侧的登记单id
	RegisterId string `json:"register_id,omitempty" xml:"register_id,omitempty"`
	// 省中文名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 市中文名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区中文名称
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 健康侧的用户支付宝id
	AlipayUserId string `json:"alipay_user_id,omitempty" xml:"alipay_user_id,omitempty"`
	// 健康侧的疫苗id
	VaccineId string `json:"vaccine_id,omitempty" xml:"vaccine_id,omitempty"`
	// 登记时间，会校验格式
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 登记人联系电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 预约渠道
	AppChannel string `json:"app_channel,omitempty" xml:"app_channel,omitempty"`
	// 省编码
	ProvinceCode int64 `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 市编码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 区编码
	AreaCode int64 `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// 健康侧的povId
	PovId int64 `json:"pov_id,omitempty" xml:"pov_id,omitempty"`
	// cdc侧的疫苗名称
	VaccineName int64 `json:"vaccine_name,omitempty" xml:"vaccine_name,omitempty"`
}
