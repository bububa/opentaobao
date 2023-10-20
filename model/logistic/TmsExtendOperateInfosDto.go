package logistic

import (
	"sync"
)

// TmsExtendOperateInfosDto 结构体
type TmsExtendOperateInfosDto struct {
	// 操作类型（枚举）： UpdateAddress-服务商修改地址
	OperateType string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
	// 操作时间(YYYY-MM-DD HH:MM:SS)
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 内容
	OperateDetail string `json:"operate_detail,omitempty" xml:"operate_detail,omitempty"`
}

var poolTmsExtendOperateInfosDto = sync.Pool{
	New: func() any {
		return new(TmsExtendOperateInfosDto)
	},
}

// GetTmsExtendOperateInfosDto() 从对象池中获取TmsExtendOperateInfosDto
func GetTmsExtendOperateInfosDto() *TmsExtendOperateInfosDto {
	return poolTmsExtendOperateInfosDto.Get().(*TmsExtendOperateInfosDto)
}

// ReleaseTmsExtendOperateInfosDto 释放TmsExtendOperateInfosDto
func ReleaseTmsExtendOperateInfosDto(v *TmsExtendOperateInfosDto) {
	v.OperateType = ""
	v.OperateTime = ""
	v.OperateDetail = ""
	poolTmsExtendOperateInfosDto.Put(v)
}
