package flight

import (
	"sync"
)

// OdInfoQueryDto 结构体
type OdInfoQueryDto struct {
	// 起飞机场
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 降落机场
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 起飞开始时间
	DepStartDate string `json:"dep_start_date,omitempty" xml:"dep_start_date,omitempty"`
	// 起飞结束时间
	DepEndDate string `json:"dep_end_date,omitempty" xml:"dep_end_date,omitempty"`
	// 0，去程；1，返程；暂时仅支持单程，默认为0
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
}

var poolOdInfoQueryDto = sync.Pool{
	New: func() any {
		return new(OdInfoQueryDto)
	},
}

// GetOdInfoQueryDto() 从对象池中获取OdInfoQueryDto
func GetOdInfoQueryDto() *OdInfoQueryDto {
	return poolOdInfoQueryDto.Get().(*OdInfoQueryDto)
}

// ReleaseOdInfoQueryDto 释放OdInfoQueryDto
func ReleaseOdInfoQueryDto(v *OdInfoQueryDto) {
	v.DepAirport = ""
	v.ArrAirport = ""
	v.DepStartDate = ""
	v.DepEndDate = ""
	v.Index = 0
	poolOdInfoQueryDto.Put(v)
}
