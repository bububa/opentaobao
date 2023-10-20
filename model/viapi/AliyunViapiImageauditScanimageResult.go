package viapi

import (
	"sync"
)

// AliyunViapiImageauditScanimageResult 结构体
type AliyunViapiImageauditScanimageResult struct {
	// 单张图片的检测结果
	SubResults []SubResult `json:"sub_results,omitempty" xml:"sub_results>sub_result,omitempty"`
	// 数据ID
	DataId string `json:"data_id,omitempty" xml:"data_id,omitempty"`
	// 任务ID
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 图像的URL
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

var poolAliyunViapiImageauditScanimageResult = sync.Pool{
	New: func() any {
		return new(AliyunViapiImageauditScanimageResult)
	},
}

// GetAliyunViapiImageauditScanimageResult() 从对象池中获取AliyunViapiImageauditScanimageResult
func GetAliyunViapiImageauditScanimageResult() *AliyunViapiImageauditScanimageResult {
	return poolAliyunViapiImageauditScanimageResult.Get().(*AliyunViapiImageauditScanimageResult)
}

// ReleaseAliyunViapiImageauditScanimageResult 释放AliyunViapiImageauditScanimageResult
func ReleaseAliyunViapiImageauditScanimageResult(v *AliyunViapiImageauditScanimageResult) {
	v.SubResults = v.SubResults[:0]
	v.DataId = ""
	v.TaskId = ""
	v.ImageUrl = ""
	poolAliyunViapiImageauditScanimageResult.Put(v)
}
