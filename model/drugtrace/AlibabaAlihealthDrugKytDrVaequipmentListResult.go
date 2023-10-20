package drugtrace

// AlibabaalihealthdrugkytdrvaequipmentlistResult 结构体
type AlibabaalihealthdrugkytdrvaequipmentlistResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 修改时间
	ModDate string `json:"mod_date,omitempty" xml:"mod_date,omitempty"`
	// 设备地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty" xml:"comments,omitempty"`
	// 设备编号
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 记录时间
	CrtDate string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
	// 设备类型
	EquipmentType string `json:"equipment_type,omitempty" xml:"equipment_type,omitempty"`
	// 关联状态
	AssociatedStatus string `json:"associated_status,omitempty" xml:"associated_status,omitempty"`
	// 关联探头
	Probe string `json:"probe,omitempty" xml:"probe,omitempty"`
	// 记录修改人
	ModIcCode string `json:"mod_ic_code,omitempty" xml:"mod_ic_code,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 地区
	DictRegionDTO string `json:"dict_region_d_t_o,omitempty" xml:"dict_region_d_t_o,omitempty"`
	// 记录创建人
	CrtIcCode string `json:"crt_ic_code,omitempty" xml:"crt_ic_code,omitempty"`
	// 所属企业
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 设备名称
	EquipmentName string `json:"equipment_name,omitempty" xml:"equipment_name,omitempty"`
	// 设备主键
	VaEquipmentId string `json:"va_equipment_id,omitempty" xml:"va_equipment_id,omitempty"`
	// 设备状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// model
	Model *AlibabaalihealthdrugkytdrvaequipmentlistModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
