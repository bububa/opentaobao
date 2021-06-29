package viapi

// AliyunViapiImageauditScanimageResult 
type AliyunViapiImageauditScanimageResult struct {
    // 数据ID
    DataId   string `json:"data_id,omitempty" xml:"data_id,omitempty"`
    // 任务ID
    TaskId   string `json:"task_id,omitempty" xml:"task_id,omitempty"`
    // 图像的URL
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    // 单张图片的检测结果
    SubResults   []SubResult `json:"sub_results,omitempty" xml:"sub_results>sub_result,omitempty"`
}
