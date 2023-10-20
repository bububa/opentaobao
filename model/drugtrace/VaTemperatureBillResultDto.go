package drugtrace

import (
	"sync"
)

// VaTemperatureBillResultDto 结构体
type VaTemperatureBillResultDto struct {
	// 最小温度值
	MinValue string `json:"min_value,omitempty" xml:"min_value,omitempty"`
	// 设备编号
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 最大温度值
	MaxValue string `json:"max_value,omitempty" xml:"max_value,omitempty"`
	// 上传企业名称
	UploadEntName string `json:"upload_ent_name,omitempty" xml:"upload_ent_name,omitempty"`
	// 设备名称
	EquipmentName string `json:"equipment_name,omitempty" xml:"equipment_name,omitempty"`
	// 单据编号
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 温度类型：存储温度/运输温度u
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolVaTemperatureBillResultDto = sync.Pool{
	New: func() any {
		return new(VaTemperatureBillResultDto)
	},
}

// GetVaTemperatureBillResultDto() 从对象池中获取VaTemperatureBillResultDto
func GetVaTemperatureBillResultDto() *VaTemperatureBillResultDto {
	return poolVaTemperatureBillResultDto.Get().(*VaTemperatureBillResultDto)
}

// ReleaseVaTemperatureBillResultDto 释放VaTemperatureBillResultDto
func ReleaseVaTemperatureBillResultDto(v *VaTemperatureBillResultDto) {
	v.MinValue = ""
	v.EquipmentCode = ""
	v.MaxValue = ""
	v.UploadEntName = ""
	v.EquipmentName = ""
	v.BillCode = ""
	v.Time = ""
	v.Type = ""
	poolVaTemperatureBillResultDto.Put(v)
}
