package dt

import (
	"sync"
)

// ExternalTaskDataImportDto 结构体
type ExternalTaskDataImportDto struct {
	// LABEL/MINE/TRAIL
	RecordType string `json:"record_type,omitempty" xml:"record_type,omitempty"`
	// 导入时的数据来源
	DataType string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 导入状态
	JobStatus string `json:"job_status,omitempty" xml:"job_status,omitempty"`
	// 导入时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolExternalTaskDataImportDto = sync.Pool{
	New: func() any {
		return new(ExternalTaskDataImportDto)
	},
}

// GetExternalTaskDataImportDto() 从对象池中获取ExternalTaskDataImportDto
func GetExternalTaskDataImportDto() *ExternalTaskDataImportDto {
	return poolExternalTaskDataImportDto.Get().(*ExternalTaskDataImportDto)
}

// ReleaseExternalTaskDataImportDto 释放ExternalTaskDataImportDto
func ReleaseExternalTaskDataImportDto(v *ExternalTaskDataImportDto) {
	v.RecordType = ""
	v.DataType = ""
	v.FileName = ""
	v.JobStatus = ""
	v.GmtCreate = ""
	v.TaskId = 0
	poolExternalTaskDataImportDto.Put(v)
}
