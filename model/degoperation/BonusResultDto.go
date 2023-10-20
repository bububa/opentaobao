package degoperation

import (
	"sync"
)

// BonusResultDto 结构体
type BonusResultDto struct {
	// updateAddress=是否填写了收货地址
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// error
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBonusResultDto = sync.Pool{
	New: func() any {
		return new(BonusResultDto)
	},
}

// GetBonusResultDto() 从对象池中获取BonusResultDto
func GetBonusResultDto() *BonusResultDto {
	return poolBonusResultDto.Get().(*BonusResultDto)
}

// ReleaseBonusResultDto 释放BonusResultDto
func ReleaseBonusResultDto(v *BonusResultDto) {
	v.Data = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Error = false
	v.Success = false
	poolBonusResultDto.Put(v)
}
