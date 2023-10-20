package omniorder

import (
	"sync"
)

// SdtStationDto 结构体
type SdtStationDto struct {
	// 站点操作时间
	ActionTime string `json:"action_time,omitempty" xml:"action_time,omitempty"`
	// 快递公司cpcode
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 快递公司名称
	CpName string `json:"cp_name,omitempty" xml:"cp_name,omitempty"`
	// 站点code
	StationCode string `json:"station_code,omitempty" xml:"station_code,omitempty"`
	// 站点联系方式
	StationContact string `json:"station_contact,omitempty" xml:"station_contact,omitempty"`
	// 站点负责人
	StationMaster string `json:"station_master,omitempty" xml:"station_master,omitempty"`
	// 站点名
	StationName string `json:"station_name,omitempty" xml:"station_name,omitempty"`
	// 站点类别（推荐站点、派送站点、揽收站点）
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolSdtStationDto = sync.Pool{
	New: func() any {
		return new(SdtStationDto)
	},
}

// GetSdtStationDto() 从对象池中获取SdtStationDto
func GetSdtStationDto() *SdtStationDto {
	return poolSdtStationDto.Get().(*SdtStationDto)
}

// ReleaseSdtStationDto 释放SdtStationDto
func ReleaseSdtStationDto(v *SdtStationDto) {
	v.ActionTime = ""
	v.CpCode = ""
	v.CpName = ""
	v.StationCode = ""
	v.StationContact = ""
	v.StationMaster = ""
	v.StationName = ""
	v.Type = ""
	poolSdtStationDto.Put(v)
}
