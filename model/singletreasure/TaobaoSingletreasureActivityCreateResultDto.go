package singletreasure

// TaobaoSingletreasureActivityCreateResultDTO 
type TaobaoSingletreasureActivityCreateResultDTO struct {
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 新建套餐 id
    Data   int64 `json:"data,omitempty" xml:"data,omitempty"`
    // 错误编码
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 系统执行成功与否
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
