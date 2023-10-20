package singletreasure

// TaobaoSingletreasureActivityItemBatchaddResultDto 结构体
type TaobaoSingletreasureActivityItemBatchaddResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *ItemProcessErrorResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
