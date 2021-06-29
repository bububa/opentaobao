package kclub

// Paging 
type Paging struct {
    // 行数
    RowCount   int64 `json:"row_count,omitempty" xml:"row_count,omitempty"`
    // 每页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 当前页
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 数据
    DataList   []KcSearchQuestionDTO `json:"data_list,omitempty" xml:"data_list>kc_search_question_dto,omitempty"`
}
