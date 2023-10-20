package lstpos

import (
	"sync"
)

// AlibabaLstPosOpenInventoryGetinventorydataResultDto 结构体
type AlibabaLstPosOpenInventoryGetinventorydataResultDto struct {
	// 接口具体返回的业务数据对象
	ModuleList []InventoryDto `json:"module_list,omitempty" xml:"module_list>inventory_dto,omitempty"`
	// 错误信息描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 业务错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 接口调用是否成功 true:调用成功 false:调用失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstPosOpenInventoryGetinventorydataResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenInventoryGetinventorydataResultDto)
	},
}

// GetAlibabaLstPosOpenInventoryGetinventorydataResultDto() 从对象池中获取AlibabaLstPosOpenInventoryGetinventorydataResultDto
func GetAlibabaLstPosOpenInventoryGetinventorydataResultDto() *AlibabaLstPosOpenInventoryGetinventorydataResultDto {
	return poolAlibabaLstPosOpenInventoryGetinventorydataResultDto.Get().(*AlibabaLstPosOpenInventoryGetinventorydataResultDto)
}

// ReleaseAlibabaLstPosOpenInventoryGetinventorydataResultDto 释放AlibabaLstPosOpenInventoryGetinventorydataResultDto
func ReleaseAlibabaLstPosOpenInventoryGetinventorydataResultDto(v *AlibabaLstPosOpenInventoryGetinventorydataResultDto) {
	v.ModuleList = v.ModuleList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaLstPosOpenInventoryGetinventorydataResultDto.Put(v)
}
