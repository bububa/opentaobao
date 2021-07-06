package alihealthmedical

// ThirdAgencyUpDownShelfRequest 结构体
type ThirdAgencyUpDownShelfRequest struct {
	// 医生外部id
	DoctorUuid string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
	// 互联网医院id
	NetHospitalId string `json:"net_hospital_id,omitempty" xml:"net_hospital_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 上下架状态：0上架，-1下架
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
