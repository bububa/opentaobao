package kclub

// Paging 
/* model for simplify = false
type Paging struct {

    // 行数
    
    RowCount   int64 `json:"row_count,omitempty"`
    

    // 每页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 当前页
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 数据
    
    DataList  struct {
        KcSearchQuestionDto  []KcSearchQuestionDto `json:"kc_search_question_dto,omitempty"`
    } `json:"data_list,omitempty"`
    

}
*/

// Paging 
type Paging struct {

    // 行数
    RowCount   int64 `json:"row_count,omitempty"`

    // 每页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 当前页
    PageNo   int64 `json:"page_no,omitempty"`

    // 数据
    DataList   []KcSearchQuestionDto `json:"data_list,omitempty"`

}
