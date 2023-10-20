package security

import (
	"sync"
)

// JaqImageDetectResult 结构体
type JaqImageDetectResult struct {
	// 异步请求任务状态；TaskProcessSuccess - 图片异步api请求task状态, 处理成功 TaskProcessing - 图片异步api请求task状态, 正在处理中 TaskInvalid - 图片异步api请求task状态, 无效task或者task不存在 TaskProcessFailed - 图片异步api请求task状态, 处理失败
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// taskId
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 图像字符识别结果结构体
	JaqOcrImageDetectResult *JaqOcrImageDetectResult `json:"jaq_ocr_image_detect_result,omitempty" xml:"jaq_ocr_image_detect_result,omitempty"`
}

var poolJaqImageDetectResult = sync.Pool{
	New: func() any {
		return new(JaqImageDetectResult)
	},
}

// GetJaqImageDetectResult() 从对象池中获取JaqImageDetectResult
func GetJaqImageDetectResult() *JaqImageDetectResult {
	return poolJaqImageDetectResult.Get().(*JaqImageDetectResult)
}

// ReleaseJaqImageDetectResult 释放JaqImageDetectResult
func ReleaseJaqImageDetectResult(v *JaqImageDetectResult) {
	v.Status = ""
	v.TaskId = ""
	v.JaqOcrImageDetectResult = nil
	poolJaqImageDetectResult.Put(v)
}
