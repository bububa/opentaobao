package lstpos

import (
	"sync"
)

// AlibabaLstPosOpenAccountCheckissettledResultDto 结构体
type AlibabaLstPosOpenAccountCheckissettledResultDto struct {
	// 错误信息描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 接口具体返回的业务数据对象
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 业务错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 接口调用是否成功 true:调用成功 false:调用失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstPosOpenAccountCheckissettledResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenAccountCheckissettledResultDto)
	},
}

// GetAlibabaLstPosOpenAccountCheckissettledResultDto() 从对象池中获取AlibabaLstPosOpenAccountCheckissettledResultDto
func GetAlibabaLstPosOpenAccountCheckissettledResultDto() *AlibabaLstPosOpenAccountCheckissettledResultDto {
	return poolAlibabaLstPosOpenAccountCheckissettledResultDto.Get().(*AlibabaLstPosOpenAccountCheckissettledResultDto)
}

// ReleaseAlibabaLstPosOpenAccountCheckissettledResultDto 释放AlibabaLstPosOpenAccountCheckissettledResultDto
func ReleaseAlibabaLstPosOpenAccountCheckissettledResultDto(v *AlibabaLstPosOpenAccountCheckissettledResultDto) {
	v.ErrorMessage = ""
	v.Module = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaLstPosOpenAccountCheckissettledResultDto.Put(v)
}
