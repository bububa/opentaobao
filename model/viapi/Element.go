package viapi

// Element 结构体
type Element struct {
	// 单个文本的检测结果
	Results []AliyunviapiimageauditscantextResult `json:"results,omitempty" xml:"results>aliyunviapiimageauditscantext_result,omitempty"`
	// 任务Id
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}
