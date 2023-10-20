package waybill

import (
	"sync"
)

// BaseResultDto 结构体
type BaseResultDto struct {
	// 异常信息
	ErrorInfoList []ErrorInfo `json:"error_info_list,omitempty" xml:"error_info_list>error_info,omitempty"`
	// 地址可达结果
	Module *ReachableRecommendResponseDto `json:"module,omitempty" xml:"module,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBaseResultDto = sync.Pool{
	New: func() any {
		return new(BaseResultDto)
	},
}

// GetBaseResultDto() 从对象池中获取BaseResultDto
func GetBaseResultDto() *BaseResultDto {
	return poolBaseResultDto.Get().(*BaseResultDto)
}

// ReleaseBaseResultDto 释放BaseResultDto
func ReleaseBaseResultDto(v *BaseResultDto) {
	v.ErrorInfoList = v.ErrorInfoList[:0]
	v.Module = nil
	v.Success = false
	poolBaseResultDto.Put(v)
}
