package lstpos

import (
	"sync"
)

// AlibabaLstPosOpenInventorySyncinventorydataResultDto 结构体
type AlibabaLstPosOpenInventorySyncinventorydataResultDto struct {
	// 接口具体返回的业务数据对象
	ModuleList []ErrorResult `json:"module_list,omitempty" xml:"module_list>error_result,omitempty"`
	// 错误消息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 接口调用是否成功 true:调用成功 false:调用失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstPosOpenInventorySyncinventorydataResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenInventorySyncinventorydataResultDto)
	},
}

// GetAlibabaLstPosOpenInventorySyncinventorydataResultDto() 从对象池中获取AlibabaLstPosOpenInventorySyncinventorydataResultDto
func GetAlibabaLstPosOpenInventorySyncinventorydataResultDto() *AlibabaLstPosOpenInventorySyncinventorydataResultDto {
	return poolAlibabaLstPosOpenInventorySyncinventorydataResultDto.Get().(*AlibabaLstPosOpenInventorySyncinventorydataResultDto)
}

// ReleaseAlibabaLstPosOpenInventorySyncinventorydataResultDto 释放AlibabaLstPosOpenInventorySyncinventorydataResultDto
func ReleaseAlibabaLstPosOpenInventorySyncinventorydataResultDto(v *AlibabaLstPosOpenInventorySyncinventorydataResultDto) {
	v.ModuleList = v.ModuleList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaLstPosOpenInventorySyncinventorydataResultDto.Put(v)
}
