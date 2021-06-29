package seaking

// TaskResultDetailDto 
type TaskResultDetailDto struct {
    // 图片翻译结果
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // 子任务状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 子任务编号
    Idx   int64 `json:"idx,omitempty" xml:"idx,omitempty"`
}
