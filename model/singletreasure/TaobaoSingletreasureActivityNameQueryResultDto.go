package singletreasure

// TaobaoSingletreasureActivityNameQueryResultDto 
type TaobaoSingletreasureActivityNameQueryResultDto struct {
    // 请求返回描述信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // data结果
    DataList   []ActivityNameCategoryDto `json:"data_list,omitempty" xml:"data_list>activity_name_category_dto,omitempty"`
    // 请求返回码
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
