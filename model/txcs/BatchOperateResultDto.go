package txcs

import (
	"sync"
)

// BatchOperateResultDto 结构体
type BatchOperateResultDto struct {
	// 失败列表
	FailList []InvoiceInputResultDto `json:"fail_list,omitempty" xml:"fail_list>invoice_input_result_dto,omitempty"`
	// 成功列表
	SuccessList []InvoiceInputResultDto `json:"success_list,omitempty" xml:"success_list>invoice_input_result_dto,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolBatchOperateResultDto = sync.Pool{
	New: func() any {
		return new(BatchOperateResultDto)
	},
}

// GetBatchOperateResultDto() 从对象池中获取BatchOperateResultDto
func GetBatchOperateResultDto() *BatchOperateResultDto {
	return poolBatchOperateResultDto.Get().(*BatchOperateResultDto)
}

// ReleaseBatchOperateResultDto 释放BatchOperateResultDto
func ReleaseBatchOperateResultDto(v *BatchOperateResultDto) {
	v.FailList = v.FailList[:0]
	v.SuccessList = v.SuccessList[:0]
	v.Status = ""
	poolBatchOperateResultDto.Put(v)
}
