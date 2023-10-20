package lstpos

import (
	"sync"
)

// AlibabaLstPosOpenCashierSynccashierdataResultDto 结构体
type AlibabaLstPosOpenCashierSynccashierdataResultDto struct {
	// 单个订单处理结果
	ModuleList []ErrorResult `json:"module_list,omitempty" xml:"module_list>error_result,omitempty"`
	// 错误消息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 调用成功与否标示 true：成功 false：失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstPosOpenCashierSynccashierdataResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenCashierSynccashierdataResultDto)
	},
}

// GetAlibabaLstPosOpenCashierSynccashierdataResultDto() 从对象池中获取AlibabaLstPosOpenCashierSynccashierdataResultDto
func GetAlibabaLstPosOpenCashierSynccashierdataResultDto() *AlibabaLstPosOpenCashierSynccashierdataResultDto {
	return poolAlibabaLstPosOpenCashierSynccashierdataResultDto.Get().(*AlibabaLstPosOpenCashierSynccashierdataResultDto)
}

// ReleaseAlibabaLstPosOpenCashierSynccashierdataResultDto 释放AlibabaLstPosOpenCashierSynccashierdataResultDto
func ReleaseAlibabaLstPosOpenCashierSynccashierdataResultDto(v *AlibabaLstPosOpenCashierSynccashierdataResultDto) {
	v.ModuleList = v.ModuleList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaLstPosOpenCashierSynccashierdataResultDto.Put(v)
}
