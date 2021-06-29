package singletreasure

// TaobaoSingletreasureActivityItemBatchaddResultDTO 
type TaobaoSingletreasureActivityItemBatchaddResultDTO struct {
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // data
    Data   *ItemProcessErrorResultDTO `json:"data,omitempty" xml:"data,omitempty"`
    // code
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
